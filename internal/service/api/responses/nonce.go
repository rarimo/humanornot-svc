package responses

import (
	"gitlab.com/rarimo/identity/kyc-service/internal/crypto"
	"gitlab.com/rarimo/identity/kyc-service/resources"
)

func NewNonce(nonce string) resources.NonceResponse {
	return resources.NonceResponse{
		Data: resources.Nonce{
			Key: resources.Key{
				Type: resources.NONCE,
			},
			Attributes: resources.NonceAttributes{
				Message: crypto.NonceToSignMessage(nonce),
			},
		},
	}
}
