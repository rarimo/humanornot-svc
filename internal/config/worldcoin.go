package config

import (
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
)

// worldcoinSettingsYamlKey is a key for the Worldcoin identity provider settings in the YAML config
const worldcoinSettingsYamlKey = "worldcoin"

// WorldcoinSettings is a struct that contains the settings for the Worldcoin identity provider
type WorldcoinSettings struct {
	BaseURL string `fig:"base_url"`
}

func (c *config) WorldcoinSettings() *WorldcoinSettings {
	return c.worldcoinSettings.Do(func() interface{} {
		cfg := WorldcoinSettings{}
		err := figure.
			Out(&cfg).
			From(kv.MustGetStringMap(c.getter, worldcoinSettingsYamlKey)).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out"))
		}

		return &cfg
	}).(*WorldcoinSettings)
}
