package responses

import (
	"github.com/rarimo/kyc-service/internal/crypto"
	"github.com/rarimo/kyc-service/resources"
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
