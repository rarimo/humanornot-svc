package responses

import (
	"github.com/google/uuid"

	"gitlab.com/rarimo/identity/kyc-service/resources"
)

type VerifyIDResponse struct {
	Data VerifyID `json:"data"`
}

type VerifyID struct {
	resources.Key
}

func NewVerify(id uuid.UUID) VerifyIDResponse {
	return VerifyIDResponse{
		Data: VerifyID{
			Key: resources.Key{
				ID:   id.String(),
				Type: resources.VERIFICATION_ID,
			},
		},
	}
}
