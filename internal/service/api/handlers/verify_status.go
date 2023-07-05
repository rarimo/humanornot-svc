package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"

	"gitlab.com/rarimo/identity/kyc-service/internal/service/api/requests"
	"gitlab.com/rarimo/identity/kyc-service/internal/service/api/responses"
)

func GetVerifyStatus(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewVerifyStatusRequest(r)
	if err != nil {
		Log(r).WithError(err).Error("failed to parse request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	user, err := KYCService(r).GetVerifyStatus(req)
	switch {
	case err != nil:
		Log(r).WithError(err).Error("failed to get user")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, responses.NewVerifyStatus(user))
}
