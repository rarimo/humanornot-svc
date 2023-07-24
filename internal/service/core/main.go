package core

import (
	"context"
	"time"

	"gitlab.com/rarimo/identity/kyc-service/internal/service/core/identity_providers/civic"
	gcpsp "gitlab.com/rarimo/identity/kyc-service/internal/service/core/identity_providers/gitcoin_passport"
	"gitlab.com/rarimo/identity/kyc-service/internal/service/core/identity_providers/worldcoin"

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
	New() KYCService

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

	isr := issuer.New(cfg.Log(), cfg.Issuer())

	return &kycService{
		db:            pg.NewMasterQ(cfg.DB()),
		issuer:        isr,
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
				ctx, isr,
			),
			providers.CivicIdentityProvider: civicIdentityProvider,
		},
	}, nil
}

func (k *kycService) New() KYCService {
	k.db = k.db.New()
	return k
}
