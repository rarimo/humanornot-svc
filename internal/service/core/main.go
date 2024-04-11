package core

import (
	"context"
	"time"

	"github.com/rarimo/humanornot-svc/internal/service/core/identity_providers/kleros"

	"github.com/rarimo/humanornot-svc/internal/service/core/identity_providers/civic"
	gcpsp "github.com/rarimo/humanornot-svc/internal/service/core/identity_providers/gitcoin_passport"
	"github.com/rarimo/humanornot-svc/internal/service/core/identity_providers/worldcoin"

	"github.com/pkg/errors"

	"github.com/rarimo/humanornot-svc/internal/config"
	"github.com/rarimo/humanornot-svc/internal/data"
	"github.com/rarimo/humanornot-svc/internal/data/pg"
	"github.com/rarimo/humanornot-svc/internal/service/api/requests"
	providers "github.com/rarimo/humanornot-svc/internal/service/core/identity_providers"
	unstopdom "github.com/rarimo/humanornot-svc/internal/service/core/identity_providers/unstoppable_domains"
	"github.com/rarimo/humanornot-svc/internal/service/core/issuer"
)

type HumanornotSvc interface {
	New() HumanornotSvc

	NewVerifyRequest(*requests.VerifyRequest) (*data.User, error)
	NewNonce(*requests.NonceRequest) (*data.Nonce, error)
	GetVerifyStatus(*requests.VerifyStatusRequest) (*data.User, error)
	GetProviderByIdentityId(*requests.GetProviderByIdentityIdRequest) (providers.IdentityProviderName, error)
}

type humanornotSvc struct {
	db                data.MasterQ
	issuer            issuer.Issuer
	identityProviders map[providers.IdentityProviderName]providers.IdentityProvider

	nonceLifetime time.Duration
}

func NewHumanornotSvc(cfg config.Config, ctx context.Context) (HumanornotSvc, error) {
	civicIdentityProvider, err := civic.NewIdentityProvider(
		cfg.Log().WithField("provider", providers.CivicIdentityProvider),
		pg.NewMasterQ(cfg.DB()),
		cfg.Civic(),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create Civic identity provider")
	}

	klerosIdentityProvider, err := kleros.NewIdentityProvider(
		cfg.Log().WithField("provider", providers.KlerosIdentityProvider),
		pg.NewMasterQ(cfg.DB()),
		cfg.Kleros(),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create Kleros identity provider")
	}

	isr := issuer.New(cfg.Log(), cfg.Issuer())

	return &humanornotSvc{
		db:            pg.NewMasterQ(cfg.DB()),
		issuer:        isr,
		nonceLifetime: cfg.HumanornotSvc().NonceLifeTime,
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
			providers.CivicIdentityProvider:  civicIdentityProvider,
			providers.KlerosIdentityProvider: klerosIdentityProvider,
		},
	}, nil
}

func (k *humanornotSvc) New() HumanornotSvc {
	k.db = k.db.New()
	return k
}
