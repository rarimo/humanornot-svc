package core

import "github.com/pkg/errors"

var (
	// Conflict errors
	ErrUserAlreadyVerifiedByEthAddress = errors.New("user already verified by eth address")
	ErrUserAlreadyVerifiedByIdentityID = errors.New("user already verified by identity id")
)
