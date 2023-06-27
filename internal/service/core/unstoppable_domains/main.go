package unstoppabledomains

import (
	"encoding/hex"
	"encoding/json"
	"net/url"

	"github.com/ethereum/go-ethereum/common"
	"github.com/imroc/req/v3"
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/logan/v3"

	"gitlab.com/rarimo/identity/kyc-service/internal/config"
	"gitlab.com/rarimo/identity/kyc-service/internal/data"
	"gitlab.com/rarimo/identity/kyc-service/internal/service/core"
)

type UnstoppableDomains struct {
	log *logan.Entry
	*req.Client
}

func New(log *logan.Entry, config *config.UnstoppableDomains) *UnstoppableDomains {
	authBaseURL := config.AuthBaseURL
	if authBaseURL.String() == "" {
		log.Debugf("Base URL for Unstoppable domains not found, the default %s is set", defaultAuthBaseURL)
		authBaseURL, _ = url.Parse(defaultAuthBaseURL)
	}

	return &UnstoppableDomains{
		log: log.WithField("identity-provider", core.UnstoppableDomainsIdentityProvider),

		Client: req.C().
			SetBaseURL(authBaseURL.String()).
			SetLogger(log),
	}
}

func (u *UnstoppableDomains) Verify(user *data.User, verifyDataRaw []byte) error {
	var verifyData VerificationData
	if err := json.Unmarshal(verifyDataRaw, &verifyData); err != nil {
		return errors.Wrap(err, "failed to unmarshal verification data")
	}

	userInfo, err := u.retrieveUserInfo(verifyData.AccessToken)
	if err != nil {
		return errors.Wrap(err, "failed to retrieve user info")
	}

	isValidSignature, err := verifyUserInfoSignature(userInfo)
	if err != nil {
		return errors.Wrap(err, "failed to verify user's signature")
	}
	if !isValidSignature {
		return ErrInvalidUsersSignature
	}

	user.EthAddress = common.HexToAddress(userInfo.WalletAddress)
	user.Status = data.UserStatusVerified

	domainInfoRaw, err := json.Marshal(Domain{
		Domain: userInfo.Domain,
	})
	if err != nil {
		return errors.Wrap(err, "failed to marshal provider data")
	}

	user.ProviderData = domainInfoRaw

	return nil
}

func (u *UnstoppableDomains) retrieveUserInfo(accessToken string) (*UserInfo, error) {
	var result UserInfo

	_, err := u.R().
		SetBearerAuthToken(accessToken).
		SetSuccessResult(&result).
		Get(userInfoEndpoint)
	if err != nil {
		return nil, errors.Wrap(err, "failed to send get request")
	}

	return &result, nil
}

func verifyUserInfoSignature(userInfo *UserInfo) (bool, error) {
	signatureBytes, err := hex.DecodeString(userInfo.EIP4361Signature)
	if err != nil {
		return false, errors.Wrap(err, "failed to decode signature from HEX")
	}

	if ok := common.IsHexAddress(userInfo.WalletAddress); !ok {
		return false, ErrInvalidWalletAddress
	}

	result, err := verifyEIP191Signature(
		signatureBytes, []byte(userInfo.EIP4361Message), common.HexToAddress(userInfo.WalletAddress),
	)
	if err != nil {
		return false, errors.Wrap(err, "failed to verify EIP191 signature")
	}

	return result, nil
}
