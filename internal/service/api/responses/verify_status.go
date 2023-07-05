package responses

import (
	"gitlab.com/rarimo/identity/kyc-service/internal/data"
	"gitlab.com/rarimo/identity/kyc-service/resources"
)

func NewVerifyStatus(user *data.User) resources.VerifyStatusResponse {
	return resources.VerifyStatusResponse{
		Data: resources.VerifyStatus{
			Key: resources.Key{
				ID:   user.ID.String(),
				Type: resources.VERIFICATION_ID,
			},
			Attributes: resources.VerifyStatusAttributes{
				Status: user.Status.String(),
			},
		},
	}
}
