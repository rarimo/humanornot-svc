package nonceCleaner

import (
	"context"
	"gitlab.com/rarimo/identity/kyc-service/internal/config"
	"gitlab.com/rarimo/identity/kyc-service/internal/data"
	"gitlab.com/rarimo/identity/kyc-service/internal/data/pg"
	"time"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/running"
)

func NewNonceCleaner(cfg config.Config) *NonceCleaner {
	return &NonceCleaner{
		storage: pg.NewMasterQ(cfg.DB()).NonceQ(),
		logger:  cfg.Log(),
	}
}

type NonceCleaner struct {
	storage data.NonceQ
	logger  *logan.Entry
}

func (s *NonceCleaner) Run(ctx context.Context) {
	running.WithBackOff(ctx,
		s.logger,
		"nonce-cleaner",
		func(ctx context.Context) error {
			return s.storage.FilterExpired().Delete()
		},
		12*time.Hour,
		1*time.Second,
		5*time.Second,
	)
}

func Run(ctx context.Context, cfg config.Config) {
	NewNonceCleaner(cfg).Run(ctx)
}
