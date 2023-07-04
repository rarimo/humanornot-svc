package handlers

import (
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/rarimo/identity/kyc-service/internal/service/api/requests"
	"gitlab.com/rarimo/identity/kyc-service/resources"
	"net/http"
)

func GetVerifyStatus(w http.ResponseWriter, r *http.Request) {
	id, err := requests.NewVerifyStatusRequest(r)
	if err != nil {
		Log(r).WithError(err).Error("failed to parse request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	user, err := MasterQueryer(r).UsersQ().WhereID(id).Get()
	if err != nil {
		Log(r).WithError(err).Error("failed to get user")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if user == nil {
		ape.RenderErr(w, problems.NotFound())
		return
	}

	ape.Render(w, resources.VerifyStatusResponse{
		Data: resources.VerifyStatus{
			Key: resources.Key{
				ID:   user.ID.String(),
				Type: resources.VERIFY_STATUS,
			},
			Attributes: resources.VerifyStatusAttributes{
				Status: string(user.Status),
			},
		},
	})
}
