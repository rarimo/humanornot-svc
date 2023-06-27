package unstoppabledomains

import (
	"context"
	"net/url"

	"github.com/imroc/req/v3"
	"gitlab.com/distributed_lab/logan/v3"

	"gitlab.com/rarimo/identity/kyc-service/internal/config"
)

type UnstoppableDomains struct {
	log *logan.Entry
	*req.Client
}

func New(log *logan.Entry, config config.UnstoppableDomains) *UnstoppableDomains {
	authBaseURL := config.AuthBaseURL
	if authBaseURL.String() == "" {
		log.Debugf("Base URL for Unstoppable domains not found, the default is set: %s", DefaultAuthBaseURL)
		authBaseURL, _ = url.Parse(DefaultAuthBaseURL)
	}

	return &UnstoppableDomains{
		log: log,

		Client: req.C().
			SetBaseURL(authBaseURL.String()).
			SetLogger(log),
	}
}

func Check(ctx context.Context) {}
