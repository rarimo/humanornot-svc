package responses

import (
	"github.com/rarimo/humanornot-svc/internal/data"
	"github.com/rarimo/humanornot-svc/resources"
)

func NewVerifyStatus(user *data.User) resources.VerifyStatusRequest {
	return resources.VerifyStatusRequest{
		Data: resources.VerifyStatus{
			Key: resources.Key{
				ID:   user.ID.String(),
				Type: resources.VERIFICATION_ID,
			},
			Attributes: resources.VerifyStatusAttributes{
				Status:  user.Status.String(),
				ClaimId: user.ClaimID.String(),
			},
		},
	}
}
