package handlers

import (
	"net/http"

	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"

	"gitlab.com/rarimo/identity/kyc-service/internal/data"
	"gitlab.com/rarimo/identity/kyc-service/internal/service/api/requests"
	"gitlab.com/rarimo/identity/kyc-service/internal/service/api/responses"
	unstoppabledomains "gitlab.com/rarimo/identity/kyc-service/internal/service/core/unstoppable_domains"
)

func UDVerify(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewVerify(r)
	if err != nil {
		Log(r).WithField("reason", err).Debug("Bad request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	user, err := KYCService(r).NewVerifyRequest(req)
	switch {
	case errors.Is(err, unstoppabledomains.ErrInvalidUsersSignature):
		Log(r).WithField("reason", err).
			WithField("identity-provider", req.IdentityProviderName).
			WithField("provider-data", string(req.ProviderData)).
			Debug("Unauthorized")
		ape.RenderErr(w, problems.Unauthorized())
	case err != nil:
		Log(r).WithError(err).
			WithField("identity-provider", req.IdentityProviderName).
			WithField("provider-data", string(req.ProviderData)).
			Error("Failed to init new verification")
		ape.RenderErr(w, problems.InternalError())
		return
	case user.Status == data.UserStatusVerified:
		w.WriteHeader(http.StatusNoContent)
		return
	}

	ape.Render(w, responses.NewVerify(user.ID))
}
