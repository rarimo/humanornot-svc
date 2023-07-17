package core

import "github.com/pkg/errors"

var (
	// Not found errors
	ErrUserNotFound = errors.New("user not found")

	// Conflict errors
	ErrUserAlreadyVerifiedByEthAddress = errors.New("eth address already taken")
	ErrUserAlreadyVerifiedByIdentityID = errors.New("identity already verified")
	ErrDuplicatedProviderData          = errors.New("duplicate provider data")
)
