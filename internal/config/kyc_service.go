package config

import (
	"time"

	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
)

type HumanornotSvc struct {
	NonceLifeTime time.Duration `fig:"nonce_life_time,required"`
}

func (c *config) HumanornotSvc() *HumanornotSvc {
	return c.humanOrNotSvc.Do(func() interface{} {
		cfg := HumanornotSvc{}
		err := figure.
			Out(&cfg).
			From(kv.MustGetStringMap(c.getter, "humanornot_svc")).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out"))
		}

		return &cfg
	}).(*HumanornotSvc)
}
