package core

import "github.com/pkg/errors"

var (
	// Not found errors
	ErrUserNotFound = errors.New("user not found")

	// Conflict errors
	ErrUserAlreadyVerifiedByEthAddress = errors.New("user already verified by eth address")
	ErrUserAlreadyVerifiedByIdentityID = errors.New("user already verified by identity id")
	ErrDuplicatedProviderData          = errors.New("duplicated provider data")
)
