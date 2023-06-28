package handlers

import (
	"net/http"

	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"

	"gitlab.com/rarimo/identity/kyc-service/internal/data"
	"gitlab.com/rarimo/identity/kyc-service/internal/service/api/requests"
	"gitlab.com/rarimo/identity/kyc-service/internal/service/api/responses"
	"gitlab.com/rarimo/identity/kyc-service/internal/service/core"
	unstopdom "gitlab.com/rarimo/identity/kyc-service/internal/service/core/identity_providers/unstoppable_domains"
)

func Verify(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewVerify(r)
	if err != nil {
		Log(r).WithField("reason", err).Debug("Bad request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	user, err := KYCService(r).NewVerifyRequest(req)
	switch {
	case errors.Is(err, core.ErrUserAlreadyVerifiedByEthAddress),
		errors.Is(err, core.ErrUserAlreadyVerifiedByIdentityID):
		Log(r).WithField("reason", err).Debug("Conflict")
		ape.RenderErr(w, problems.Conflict())
		return
	case errors.Is(err, unstopdom.ErrInvalidUsersSignature),
		errors.Is(err, unstopdom.ErrInvalidAccessToken):
		Log(r).WithField("reason", err).
			WithField("identity-provider", req.IdentityProviderName).
			WithField("provider-data", string(req.ProviderData)).
			Debug("Unauthorized")
		ape.RenderErr(w, problems.Unauthorized())
		return
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
