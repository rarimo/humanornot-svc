package core

import (
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"gitlab.com/rarimo/identity/kyc-service/internal/config"
	"gitlab.com/rarimo/identity/kyc-service/internal/data"
	"gitlab.com/rarimo/identity/kyc-service/internal/data/pg"
	"gitlab.com/rarimo/identity/kyc-service/internal/service/api/requests"
	unstoppabledomains "gitlab.com/rarimo/identity/kyc-service/internal/service/core/unstoppable_domains"
)

type KYCService interface {
	NewVerifyRequest(*requests.VerifyRequest) (*data.User, error)
}

type kycService struct {
	identityProviders map[IdentityProviderName]IdentityProvider
	db                data.MasterQ
	issuer            issuer.Issuer
}

func NewKYCService(cfg config.Config) KYCService {
	return &kycService{
		db: pg.NewMasterQ(cfg.DB()),
		identityProviders: map[IdentityProviderName]IdentityProvider{
			UnstoppableDomainsIdentityProvider: unstoppabledomains.New(cfg.Log(), cfg.UnstoppableDomains()),
		},
	}
}

func (k *kycService) NewVerifyRequest(req *requests.VerifyRequest) (*data.User, error) {
	newUser := data.User{
		ID:         uuid.New(),
		Status:     data.UserStatusInitialized,
		CreatedAt:  time.Now(),
		IdentityID: req.IdentityID,
	}

	err := k.identityProviders[req.IdentityProviderName].Verify(&newUser, req.ProviderData)
	if err != nil {
		return nil, errors.Wrap(err, "failed to verify user")
	}

	err = k.db.VerificationUsersQ().Insert(&newUser)
	if err != nil {
		return nil, errors.Wrap(err, "failed to insert new user into db")
	}

	return &newUser, nil
}
