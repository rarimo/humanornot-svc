package core

import "context"

type IdentityProvider interface {
	Verify(context.Context) error
}
