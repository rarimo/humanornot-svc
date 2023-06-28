package config

import (
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
)

type UnstoppableDomains struct {
	AuthBaseURL string `fig:"auth_base_url"`
}

func (c *config) UnstoppableDomains() *UnstoppableDomains {
	return c.unstoppableDomains.Do(func() interface{} {
		cfg := UnstoppableDomains{}
		err := figure.
			Out(&cfg).
			From(kv.MustGetStringMap(c.getter, "unstoppable_domains")).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out"))
		}

		return &cfg
	}).(*UnstoppableDomains)
}
