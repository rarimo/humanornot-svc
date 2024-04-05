package handlers

import (
	"net/http"

	"github.com/pkg/errors"
	"github.com/rarimo/kyc-service/internal/service/api/requests"
	identityproviders "github.com/rarimo/kyc-service/internal/service/core/identity_providers"
	"github.com/rarimo/kyc-service/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetProviderByIdentityId(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewGetProviderByIdentityIdRequest(r)
	if err != nil {
		Log(r).WithField("reason", err).Debug("Bad request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	provider, err := KYCService(r).GetProviderByIdentityId(req)
	switch {
	case errors.Is(err, identityproviders.ErrProviderNotFound), errors.Is(err, identityproviders.ErrUserNotFound):
		Log(r).WithField("reason", err).Debug("Not found")
		ape.RenderErr(w, problems.NotFound())
		return
	case err != nil:
		Log(r).WithError(err).Error("Failed to get provider by did")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	response := resources.GetProviderByIdentityIdResponse{
		Data: resources.GetProviderByIdentityId{
			Key: resources.Key{
				ID:   req.IdentityID.String(),
				Type: resources.PROVIDER,
			},
			Attributes: resources.GetProviderByIdentityIdAttributes{
				Provider: provider.String(),
			},
		},
	}

	ape.Render(w, response)
}
