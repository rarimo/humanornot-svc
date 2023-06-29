package worldcoin

import (
	"fmt"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const (
	defaultBaseURL   = "https://id.worldcoin.org"
	userInfoEndpoint = "/userinfo"

	likelyHumanStrong = "strong"
)

type (
	// VerificationData is data that is required by Worldcoin to verify a user
	VerificationData struct {
		IdToken string `json:"id_token"`
	}

	// UserInfo is data that is returned by Worldcoin by /userinfo endpoint.
	UserInfo struct {
		Sub          string       `json:"sub"`
		HumanityInfo HumanityInfo `json:"https://id.worldcoin.org/beta"`
		Email        string       `json:"email,omitempty"`
		Name         string       `json:"name,omitempty"`
		GivenName    string       `json:"given_name,omitempty"`
		FamilyName   string       `json:"family_name,omitempty"`
	}

	HumanityInfo struct {
		LikelyHuman    string `json:"likely_human"`
		CredentialType string `json:"credential_type"`
	}
)

var (
	ErrInvalidIdToken = errors.New("invalid id token")
	ErrNotLikelyHuman = errors.New("not likely human")
)

func unexpectedStatusCode(statusCode int) error {
	return errors.New(fmt.Sprintf("received unexpected status code: %d", statusCode))
}
