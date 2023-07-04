package identityproviders

import "github.com/pkg/errors"

var (
	// Unauthorized errors
	ErrNonceNotFound         = errors.New("nonce for provided address not found")
	ErrInvalidUsersSignature = errors.New("invalid signature")
	ErrInvalidAccessToken    = errors.New("invalid access token")

	// Bad request errors
	ErrInvalidVerificationData = errors.New("verification data is invalid")
)

type IdentityProviderName string

func (ipn IdentityProviderName) String() string {
	return string(ipn)
}

const (
	UnstoppableDomainsIdentityProvider IdentityProviderName = "unstoppable_domains"
	CivicIdentityProvider              IdentityProviderName = "civic"
	GitCoinPassportIdentityProvider    IdentityProviderName = "gitcoin_passport"
	WorldCoinIdentityProvider          IdentityProviderName = "worldcoin"
)

var IdentityProviderNames = map[string]IdentityProviderName{
	UnstoppableDomainsIdentityProvider.String(): UnstoppableDomainsIdentityProvider,
	CivicIdentityProvider.String():              CivicIdentityProvider,
	GitCoinPassportIdentityProvider.String():    GitCoinPassportIdentityProvider,
	WorldCoinIdentityProvider.String():          WorldCoinIdentityProvider,
}
