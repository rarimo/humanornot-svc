package responses

import (
	"gitlab.com/rarimo/identity/kyc-service/internal/crypto"
	"gitlab.com/rarimo/identity/kyc-service/resources"
)

func NewNonce(nonce string) resources.NewNonceResponse {
	return resources.NewNonceResponse{
		Data: resources.NewNonce{
			Key: resources.Key{
				Type: resources.NONCE,
			},
			Attributes: resources.NewNonceAttributes{
				Message: crypto.NonceToSignMessage(nonce),
			},
		},
	}
}
