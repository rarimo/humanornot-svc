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
	ClaimTypeNaturalPerson ClaimType = "NaturalPerson"
)

type IsNaturalPersonCredentialSubject struct {
	IsNatural string `json:"is_natural"`
}

type IssueClaimResponse struct {
	Data IssueClaimResponseData `json:"data"`
}

type IssueClaimResponseData struct {
	resources.Key
}
