package issuer

import (
	"github.com/pkg/errors"
	"gitlab.com/rarimo/identity/issuer/resources"
)

var (
	ErrUnexpectedStatusCode = errors.New("unexpected status code")
)

const (
	issueEndpoint = "/private/claims/issue/{identity_id}/{claim_type}"
)

const (
	identityIDPathParam = "identity_id"
	claimTypePathParam  = "claim_type"
)

type ClaimType string

func (c ClaimType) String() string {
	return string(c)
}

const (
	ClaimTypeNaturalPerson     ClaimType = "NaturalPerson"
	ClaimTypeIdentityProviders ClaimType = "IdentityProviders"
)

type IdentityProviderName string

func (ipn IdentityProviderName) String() string {
	return string(ipn)
}

const (
	UnstoppableDomainsProviderName IdentityProviderName = "Unstoppable Domains"
	CivicProviderName              IdentityProviderName = "Civic"
	GitcoinProviderName            IdentityProviderName = "Gitcoin Passport"
	WorldCoinProviderName          IdentityProviderName = "Worldcoin"
)

type IsNaturalPersonCredentialSubject struct {
	IsNatural string `json:"is_natural"`
}

type IdentityProvidersCredentialSubject struct {
	Provider                 IdentityProviderName `json:"provider"`
	IsNatural                bool                 `json:"is_natural"`
	Address                  string               `json:"address"`
	GitcoinPassportScore     string               `json:"gitcoin_passport_score"`
	WorldCoinScore           string               `json:"worldcoin_score"`
	UnstoppableDomain        string               `json:"unstoppable_domain"`
	CivicGatekeeperNetworkID string               `json:"civic_gatekeeper_network_id"`
	KYCAdditionalData        string               `json:"kyc_additional_data"`
}

type IssueClaimResponse struct {
	Data IssueClaimResponseData `json:"data"`
}

type IssueClaimResponseData struct {
	resources.Key
}
