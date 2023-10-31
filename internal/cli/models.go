package cli

import (
	"context"

	"github.com/rarimo/kyc-service/internal/config"
)

type (
	// RunnerFunc is a function that runs a service.
	RunnerFunc = func(context context.Context, config config.Config)
)
