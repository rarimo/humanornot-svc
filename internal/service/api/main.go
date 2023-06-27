package api

import (
	"context"
	"net"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"

	"gitlab.com/rarimo/identity/kyc-service/internal/config"
	"gitlab.com/rarimo/identity/kyc-service/internal/service/core"
)

type service struct {
	log        *logan.Entry
	copus      types.Copus
	listener   net.Listener
	kycService core.KYCService
}

func newService(cfg config.Config) *service {
	return &service{
		log:        cfg.Log(),
		copus:      cfg.Copus(),
		listener:   cfg.Listener(),
		kycService: core.NewKYCService(cfg),
	}
}

func Run(ctx context.Context, cfg config.Config) {
	svc := newService(cfg)

	if err := svc.copus.RegisterChi(svc.router()); err != nil {
		panic(errors.Wrap(err, "cop failed"))
	}

	svc.log.Info("Service started")
	ape.Serve(ctx, svc.router(), cfg, ape.ServeOpts{})
}
