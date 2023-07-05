package core

import (
	"context"
	"time"

	"gitlab.com/rarimo/identity/kyc-service/internal/crypto"
	"gitlab.com/rarimo/identity/kyc-service/internal/service/core/identity_providers/civic"
	gcpsp "gitlab.com/rarimo/identity/kyc-service/internal/service/core/identity_providers/gitcoin_passport"
	"gitlab.com/rarimo/identity/kyc-service/internal/service/core/identity_providers/worldcoin"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"gitlab.com/rarimo/identity/kyc-service/internal/config"
	"gitlab.com/rarimo/identity/kyc-service/internal/data"
	"gitlab.com/rarimo/identity/kyc-service/internal/data/pg"
	"gitlab.com/rarimo/identity/kyc-service/internal/service/api/requests"
	providers "gitlab.com/rarimo/identity/kyc-service/internal/service/core/identity_providers"
	unstopdom "gitlab.com/rarimo/identity/kyc-service/internal/service/core/identity_providers/unstoppable_domains"
	"gitlab.com/rarimo/identity/kyc-service/internal/service/core/issuer"
)

type KYCService interface {
	NewVerifyRequest(*requests.VerifyRequest) (*data.User, error)
	NewNonce(*requests.NonceRequest) (*data.Nonce, error)
	GetVerifyStatus(*requests.VerifyStatusRequest) (*data.User, error)
}

type kycService struct {
	db                data.MasterQ
	issuer            issuer.Issuer
	identityProviders map[providers.IdentityProviderName]providers.IdentityProvider

	nonceLifetime time.Duration
}

func NewKYCService(cfg config.Config, ctx context.Context) (KYCService, error) {
	civicIdentityProvider, err := civic.NewIdentityProvider(
		cfg.Log().WithField("provider", providers.CivicIdentityProvider),
		pg.NewMasterQ(cfg.DB()),
		cfg.Civic(),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create Civic identity provider")
	}

	return &kycService{
		db:            pg.NewMasterQ(cfg.DB()),
		issuer:        issuer.New(cfg.Log(), cfg.Issuer()),
		nonceLifetime: cfg.KYCService().NonceLifeTime,
		identityProviders: map[providers.IdentityProviderName]providers.IdentityProvider{
			providers.UnstoppableDomainsIdentityProvider: unstopdom.NewIdentityProvider(
				cfg.Log().WithField("provider", providers.UnstoppableDomainsIdentityProvider),
				cfg.UnstoppableDomains(),
			),
			providers.WorldCoinIdentityProvider: worldcoin.NewIdentityProvider(
				cfg.Log().WithField("provider", providers.WorldCoinIdentityProvider),
				cfg.WorldcoinSettings(),
			),
			providers.GitCoinPassportIdentityProvider: gcpsp.NewIdentityProvider(
				cfg.Log().WithField("provider", providers.GitCoinPassportIdentityProvider),
				cfg.GitcoinPassportSettings(),
				pg.NewMasterQ(cfg.DB()),
				ctx,
			),
			providers.CivicIdentityProvider: civicIdentityProvider,
		},
	}, nil
}

func (k *kycService) NewVerifyRequest(req *requests.VerifyRequest) (*data.User, error) {
	prevUser, err := k.db.UsersQ().WhereIdentityID(data.NewIdentityID(req.IdentityID)).Get()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user from db with the same identityID")
	}
	if prevUser != nil {
		if prevUser.Status != data.UserStatusUnverified {
			return nil, ErrUserAlreadyVerifiedByIdentityID
		}
	}

	newUser := data.User{
		ID:         uuid.New(),
		Status:     data.UserStatusInitialized,
		CreatedAt:  time.Now(),
		IdentityID: data.NewIdentityID(req.IdentityID),
	}

	if err = k.verifyUser(req, &newUser); err != nil {
		return nil, errors.Wrap(err, "failed to verify user")
	}

	err = k.db.Transaction(func() error {
		if err = k.db.UsersQ().Upsert(&newUser); err != nil {
			return errors.Wrap(err, "failed to insert new user into db")
		}

		if newUser.Status == data.UserStatusVerified {
			_, err := k.issuer.IssueClaim(
				newUser.IdentityID.ID,
				issuer.ClaimTypeNaturalPerson,
				issuer.IsNaturalPersonCredentialSubject{
					IsNatural: "1",
				})
			if err != nil {
				return errors.Wrap(err, "failed to issue claim")
			}
		}

		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute db transaction")
	}

	return &newUser, nil
}

func (k *kycService) NewNonce(req *requests.NonceRequest) (*data.Nonce, error) {
	if err := k.db.NonceQ().WhereEthAddress(req.Address).Delete(); err != nil {
		return nil, errors.Wrap(err, "failed to delete old nonce")
	}

	newNonce, err := crypto.NewNonce()
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate nonce")
	}

	nonce := &data.Nonce{
		EthAddress: req.Address,
		Nonce:      newNonce,
		ExpiresAt:  time.Now().Add(k.nonceLifetime),
	}

	err = k.db.NonceQ().Insert(nonce)
	if err != nil {
		return nil, errors.Wrap(err, "failed to insert nonce into db")
	}

	return nonce, nil
}

func (k *kycService) GetVerifyStatus(req *requests.VerifyStatusRequest) (*data.User, error) {
	user, err := k.db.UsersQ().WhereID(req.VerifyID).Get()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user from db")
	}
	if user == nil {
		return nil, ErrUserNotFound
	}

	return user, nil
}

func (k *kycService) verifyUser(req *requests.VerifyRequest, newUser *data.User) error {
	providerHash, err := k.identityProviders[req.IdentityProviderName].Verify(newUser, req.ProviderData)
	if err != nil {
		return errors.Wrap(err, "failed to verify with identity provider")
	}

	user, err := k.db.UsersQ().WhereProviderHash(providerHash).Get()
	if err != nil {
		return errors.Wrap(err, "failed to get user from db with the same providerDataHash")
	}
	if user != nil {
		return ErrDuplicatedProviderData
	}

	newUser.ProviderHash = providerHash

	return nil
}
