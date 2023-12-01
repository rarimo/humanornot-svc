package config

import (
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/copus"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/kit/pgdb"
)

type Config interface {
	comfig.Logger
	pgdb.Databaser
	types.Copuser
	comfig.Listenerer

	Civic() *Civic
	WorldcoinSettings() *WorldcoinSettings
	UnstoppableDomains() *UnstoppableDomains
	GitcoinPassportSettings() *GitcoinPassportSettings
	Kleros() *Kleros
	Issuer() *Issuer
	KYCService() *KYCService
}

type config struct {
	comfig.Logger
	pgdb.Databaser
	types.Copuser
	comfig.Listenerer
	getter kv.Getter

	civic                   comfig.Once
	unstoppableDomains      comfig.Once
	gitcoinPassportSettings comfig.Once
	worldcoinSettings       comfig.Once
	issuer                  comfig.Once
	kycService              comfig.Once
}

func New(getter kv.Getter) Config {
	return &config{
		getter:     getter,
		Databaser:  pgdb.NewDatabaser(getter),
		Copuser:    copus.NewCopuser(getter),
		Listenerer: comfig.NewListenerer(getter),
		Logger:     comfig.NewLogger(getter, comfig.LoggerOpts{}),
	}
}
