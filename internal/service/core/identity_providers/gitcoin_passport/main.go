package gcpsp

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/imroc/req/v3"
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/rarimo/identity/kyc-service/internal/config"
	"gitlab.com/rarimo/identity/kyc-service/internal/crypto"
	"gitlab.com/rarimo/identity/kyc-service/internal/data"
	"net/http"
	"strconv"
	"time"
)

// GitcoinPassport is an identity provider
type GitcoinPassport struct {
	cli          *req.Client
	logger       *logan.Entry
	settings     config.GitcoinPassportSettings
	masterQ      data.MasterQ
	scoreReqChan chan data.User
}

// NewIdentityProvider creates a new GitcoinPassport instance
func NewIdentityProvider(
	settings config.GitcoinPassportSettings,
	masterQ data.MasterQ,
	logger *logan.Entry,
	ctx context.Context,
) *GitcoinPassport {
	instance := &GitcoinPassport{
		cli:          req.C().SetBaseURL(settings.BaseURL).SetLogger(logger),
		settings:     settings,
		masterQ:      masterQ,
		scoreReqChan: make(chan data.User),
	}

	go instance.watchNewCheckScoreRequest(ctx)

	return instance
}

func (g *GitcoinPassport) Verify(user *data.User, verifyProviderDataRaw []byte) error {
	var verifyData VerificationData
	if err := json.Unmarshal(verifyProviderDataRaw, &verifyData); err != nil {
		return errors.Wrap(err, "failed to unmarshal verification data")
	}

	if err := verifyData.Validate(); err != nil {
		return errors.Wrap(err, "verification data is invalid")
	}

	userAddr := common.HexToAddress(verifyData.Address)

	if !g.settings.SkipSigCheck {
		nonce, err := g.masterQ.NonceQ().FilterByAddress(verifyData.Address).Get()
		if err != nil {
			return errors.Wrap(err, "failed to get nonce")
		}

		if nonce == nil {
			return errors.New("nonce for provided address not found")
		}

		valid, err := crypto.VerifyEIP191Signature(verifyData.Signature, nonce.Message, userAddr)
		if err != nil {
			return errors.Wrap(err, "failed to verify signature")
		}

		if !valid {
			return errors.New("signature is invalid")
		}
	}

	response, err := g.submitUserPassport(verifyData.Address)
	if err != nil {
		return errors.Wrap(err, "failed to submit user's passport")
	}

	rawData, err := json.Marshal(response.ProviderData)
	if err != nil {
		return errors.Wrap(err, "failed to marshal provider data")
	}

	user.Status = data.UserStatusPending
	user.EthAddress = &userAddr

	switch response.Status {
	case statusDone:
		if err := g.validateScore(response.Score); err != nil {
			return errors.Wrap(err, "failed to validate score")
		}

		user.ProviderData = rawData
		user.Status = data.UserStatusVerified
	case statusProcessing:
		g.scoreReqChan <- *user
	default:
		return errors.New(fmt.Sprintf("unknown response status: %s", response.Status))
	}

	return nil
}

// watchNewCheckScoreRequest watches for new requests to check user's score
func (g *GitcoinPassport) watchNewCheckScoreRequest(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case user := <-g.scoreReqChan:
			go func(user data.User) {
				if err := g.processNewCheckScoreRequest(user); err != nil {
					g.logger.WithError(err).Error("failed to process new check score request")
				}
			}(user)
		}
	}
}

// processNewCheckScoreRequest processes new request to check user's score
func (g *GitcoinPassport) processNewCheckScoreRequest(user data.User) error {
	userAddr := user.EthAddress.String()

	for i := 0; i < g.settings.GetScoreMaxRetries; i++ {
		score, processed, err := g.getUserScore(userAddr)
		if err != nil {
			return errors.Wrap(err, "failed to get user score")
		}

		if !processed {
			<-time.After(retryInterval)
			continue
		}

		if err := g.validateScore(score); err != nil {
			user.Status = data.UserStatusRejected
		}

		providerDataRaw, err := json.Marshal(ProviderData{
			Address: userAddr,
			Score:   score,
		})
		if err != nil {
			return errors.Wrap(err, "failed to marshal provider data")
		}

		user.ProviderData = providerDataRaw

		return g.masterQ.UsersQ().Update(&user)
	}

	return errors.New("max retries exceeded")
}

// getUserScore returns user's score and whether it's processed or not from Gitcoin Passport
func (g *GitcoinPassport) getUserScore(address string) (score string, processed bool, err error) {
	var resp SubmitPassportResponse

	response, err := g.cli.R().
		SetHeader("X-API-KEY", g.settings.APIkey).
		SetSuccessResult(&resp).
		Get(fmt.Sprintf("%s/%v/%s", scoreEndpoint, g.settings.ScorerId, address))

	if err != nil {
		return score, processed, errors.Wrap(err, "failed to send request")
	}

	if response.StatusCode >= 299 {
		if response.StatusCode == http.StatusUnauthorized {
			return score, processed, err
		}

		return score, processed, err
	}

	switch resp.Status {
	case statusDone:
		return resp.Score, true, nil
	case statusProcessing:
		return score, processed, nil
	default:
		return score, processed, errors.New("unknown response status")
	}
}

// submitUserPassport submits user's passport to the provider
func (g *GitcoinPassport) submitUserPassport(address string) (*SubmitPassportResponse, error) {
	raw, err := json.Marshal(SubmitPassportRequest{
		ScorerId: g.settings.ScorerId,
		Address:  address,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal request")
	}

	var resp SubmitPassportResponse

	response, err := g.cli.R().
		SetHeader("X-API-KEY", g.settings.APIkey).
		SetBody(raw).
		SetSuccessResult(&resp).
		Post(submitPassportEndpoint)
	if err != nil {
		return nil, errors.Wrap(err, "failed to send request")
	}

	if response.StatusCode >= 299 {
		if response.StatusCode == http.StatusUnauthorized {
			return nil, err
		}

		return nil, err
	}

	return &resp, nil
}

// validateScore validates score received from the provider and returns an error if it's invalid
func (g *GitcoinPassport) validateScore(score string) error {
	value, err := strconv.ParseFloat(score, 64)
	if err != nil {
		return errors.Wrap(err, "failed to parse score")
	}

	if value < g.settings.GateScore {
		return errors.New("score is too low")
	}

	return nil
}
