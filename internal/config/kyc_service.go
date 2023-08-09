package config

import (
	"time"

	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
)

type KYCService struct {
	NonceLifeTime time.Duration `fig:"nonce_life_time,required"`
}

func (c *config) KYCService() *KYCService {
	return c.kycService.Do(func() interface{} {
		cfg := KYCService{}
		err := figure.
			Out(&cfg).
			From(kv.MustGetStringMap(c.getter, "kyc_service")).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out"))
		}

		return &cfg
	}).(*KYCService)
}
