package config

import (
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
)

// gitcoinPassportSettingsYamlKey is a key for the Gitcoin Passport identity provider settings in the YAML config
const gitcoinPassportSettingsYamlKey = "gitcoin_passport"

// GitcoinPassportSettings is a struct that contains the settings for the Gitcoin Passport identity provider
type GitcoinPassportSettings struct {
	// APIkey is the API key for the Gitcoin Passport identity provider
	APIkey string `fig:"api_key,required"`
	// BaseURL is the base URL for the Gitcoin Passport identity provider
	BaseURL string `fig:"base_url"`
	// ScorerId is the ID of the Scorer instance to use for scoring
	ScorerId int `fig:"scorer_id,required"`
	// GateScore is the minimum score required to pass the verification
	GateScore float64 `fig:"gate_score,required"`
	// GetScoreMaxRetries is the maximum number of retries for the GetScore request
	GetScoreMaxRetries int `fig:"get_score_max_retries"`

	// SkipSigCheck is a flag that indicates whether to skip signature verification
	SkipSigCheck bool `fig:"skip_sig_check"`
}

// GitcoinPassportSettings returns the Gitcoin Passport identity provider settings
func (c *config) GitcoinPassportSettings() *GitcoinPassportSettings {
	return c.gitcoinPassportSettings.Do(func() interface{} {
		cfg := GitcoinPassportSettings{}

		err := figure.
			Out(&cfg).
			From(kv.MustGetStringMap(c.getter, gitcoinPassportSettingsYamlKey)).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out gitcoin passport settings"))
		}

		return &cfg
	}).(*GitcoinPassportSettings)
}
