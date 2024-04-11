package responses

import (
	"github.com/rarimo/humanornot-svc/internal/crypto"
	"github.com/rarimo/humanornot-svc/resources"
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
