package gcpsp

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	cryptoPkg "github.com/ethereum/go-ethereum/crypto"
	"github.com/imroc/req/v3"
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/logan/v3"

	"gitlab.com/rarimo/identity/kyc-service/internal/config"
	"gitlab.com/rarimo/identity/kyc-service/internal/crypto"
	"gitlab.com/rarimo/identity/kyc-service/internal/data"
	providers "gitlab.com/rarimo/identity/kyc-service/internal/service/core/identity_providers"
)

// GitcoinPassport is an identity provider
type GitcoinPassport struct {
	cli          *req.Client
	logger       *logan.Entry
	settings     *config.GitcoinPassportSettings
	masterQ      data.MasterQ
	scoreReqChan chan data.User
}

// NewIdentityProvider creates a new GitcoinPassport instance
func NewIdentityProvider(
	logger *logan.Entry,
	settings *config.GitcoinPassportSettings,
	masterQ data.MasterQ,
	ctx context.Context,
) *GitcoinPassport {
	if settings.BaseURL == "" {
		logger.Debugf("Base URL for Gitcoin Passport not found, the default %s is set", defaultBaseURL)
		settings.BaseURL = defaultBaseURL
	}

	if settings.GetScoreMaxRetries == 0 {
		logger.Debugf("GetScoreMaxRetries for Gitcoin Passport not found, the default %v is set", defaultRetryCount)
		settings.GetScoreMaxRetries = defaultRetryCount
	}

	instance := &GitcoinPassport{
		cli:          req.C().SetBaseURL(settings.BaseURL).SetLogger(logger),
		settings:     settings,
		masterQ:      masterQ,
		scoreReqChan: make(chan data.User),
	}

	go instance.watchNewCheckScoreRequest(ctx)

	return instance
}

func (g *GitcoinPassport) Verify(user *data.User, verifyDataRaw []byte) ([]byte, error) {
	var verifyData VerificationData
	if err := json.Unmarshal(verifyDataRaw, &verifyData); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal verification data")
	}

	if err := verifyData.Validate(); err != nil {
		return nil, providers.ErrInvalidVerificationData
	}

	userAddr := common.HexToAddress(verifyData.Address)

	if err := g.verifySignature(verifyData.Signature, userAddr); err != nil {
		return nil, errors.Wrap(err, "failed to verify signature")
	}

	response, err := g.submitUserPassport(verifyData.Address)
	if err != nil {
		return nil, errors.Wrap(err, "failed to submit user's passport")
	}

	rawData, err := json.Marshal(response.ProviderData)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal provider data")
	}

	user.Status = data.UserStatusPending
	user.EthAddress = &userAddr

	switch response.Status {
	case statusDone:
		if err := g.validateScore(response.Score); err != nil {
			return nil, errors.Wrap(err, "failed to validate score")
		}

		user.ProviderData = rawData
		user.Status = data.UserStatusVerified
	case statusProcessing:
		g.scoreReqChan <- *user
	default:
		return nil, errors.Wrapf(ErrUnexpectedStatus, response.Status)
	}

	return cryptoPkg.Keccak256(
		userAddr.Bytes(),
		providers.GitCoinPassportIdentityProvider.Bytes(),
	), nil
}

// verifySignature verifies user's signature
func (g *GitcoinPassport) verifySignature(signature string, userAddr common.Address) error {
	if g.settings.SkipSigCheck {
		return nil
	}

	nonce, err := g.masterQ.NonceQ().
		WhereEthAddress(userAddr).
		WhereExpiresAtGt(time.Now()).
		Get()
	if err != nil {
		return errors.Wrap(err, "failed to get nonce")
	}
	if nonce == nil {
		return providers.ErrNonceNotFound
	}

	valid, err := crypto.VerifyEIP191Signature(
		signature,
		crypto.NonceToSignMessage(nonce.Nonce),
		userAddr,
	)
	if err != nil || !valid {
		return providers.ErrInvalidUsersSignature
	}

	return errors.Wrap(g.masterQ.NonceQ().WhereEthAddress(userAddr).Delete(), "failed to delete nonce")
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

	// wait for some time before checking user's score.
	// this time sleep also waits when the user will be processed and saved in db by the main goroutine
	<-time.After(retryInterval)

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
			user.Status = data.UserStatusUnverified
		}

		providerDataRaw, err := json.Marshal(ProviderData{
			Address: userAddr,
			Score:   score,
		})
		if err != nil {
			return errors.Wrap(err, "failed to marshal provider data")
		}

		user.ProviderData = providerDataRaw
		return errors.Wrap(g.masterQ.UsersQ().Update(&user), "failed to update user")
	}

	user.Status = data.UserStatusUnverified
	return errors.Wrap(g.masterQ.UsersQ().Update(&user), "failed to update user")
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
			return score, processed, providers.ErrInvalidAccessToken
		}

		return score, processed, errors.Wrapf(ErrUnexpectedStatusCode, response.String())
	}

	switch resp.Status {
	case statusDone:
		return resp.Score, true, nil
	case statusProcessing:
		return score, processed, nil
	default:
		return score, processed, errors.Wrapf(ErrUnexpectedStatus, resp.Status)
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
			return nil, providers.ErrInvalidAccessToken
		}

		return nil, errors.Wrapf(ErrUnexpectedStatusCode, response.String())
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
		return ErrScoreIsTooLow
	}

	return nil
}
