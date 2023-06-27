package core

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
