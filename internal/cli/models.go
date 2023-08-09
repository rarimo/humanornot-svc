package cli

import (
	"context"
	"gitlab.com/rarimo/identity/kyc-service/internal/config"
)

type (
	// RunnerFunc is a function that runs a service.
	RunnerFunc = func(context context.Context, config config.Config)
)
