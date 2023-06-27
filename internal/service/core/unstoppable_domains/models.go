package unstoppabledomains

import "github.com/pkg/errors"

var (
	// Internal errors
	ErrInvalidWalletAddress = errors.New("invalid wallet address")

	// Unathorized errors
	ErrInvalidUsersSignature = errors.New("invalid signature")
)

const (
	defaultAuthBaseURL = "https://auth.unstoppabledomains.com"
	userInfoEndpoint   = "/userinfo"
)

// UserInfo data that is returned by Unstoppable Domains by /userinfo endpoint.
// It contains more information, but we only need this fields
type UserInfo struct {
	EIP4361Message   string `json:"eip4361_message"`
	EIP4361Signature string `json:"eip4361_signature"`
	WalletAddress    string `json:"wallet_address"`
	Domain           string `json:"sub"`
}

// VerificationData data that is required by Unstoppable Domains to verify a user
type VerificationData struct {
	AccessToken string `json:"access_token"`
}

type Domain struct {
	Domain string `json:"domain"`
}
