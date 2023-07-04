package identityproviders

import (
	"gitlab.com/rarimo/identity/kyc-service/internal/data"
)

type IdentityProvider interface {
	Verify(user *data.User, verifyProviderDataRaw []byte) error
}
