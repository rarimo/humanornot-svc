package config

import (
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
)

type Issuer struct {
	BaseURL      string `fig:"base_url"`
	AuthUsername string `fig:"auth_username"`
	AuthPassword string `fig:"auth_password"`
	SchemaType   string `fig:"schema_type"`
	SchemaURL    string `fig:"schema_url"`
}

func (c *config) Issuer() *Issuer {
	return c.issuer.Do(func() interface{} {
		cfg := Issuer{}
		err := figure.
			Out(&cfg).
			From(kv.MustGetStringMap(c.getter, "issuer")).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out"))
		}

		if cfg.BaseURL == "" {
			panic(errors.New("issuer.base_url is required"))
		}

		return &cfg
	}).(*Issuer)
}
