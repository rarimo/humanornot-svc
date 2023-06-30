package handlers

import (
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/rarimo/identity/kyc-service/internal/crypto"
	"gitlab.com/rarimo/identity/kyc-service/internal/data"
	"gitlab.com/rarimo/identity/kyc-service/internal/service/api/requests"
	"gitlab.com/rarimo/identity/kyc-service/resources"
	"net/http"
	"time"
)

func GetNonce(w http.ResponseWriter, r *http.Request) {
	logger := Log(r)
	db := MasterQueryer(r)

	request, err := requests.NewNonceRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	nonce, err := crypto.NewNonce()
	if err != nil {
		logger.WithError(err).Error("failed to generate nonce")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if err := db.NonceQ().FilterByAddress(request.Attributes.Address).Delete(); err != nil {
		logger.WithError(err).Error("failed to delete user nonce")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	_, err = db.NonceQ().Insert(data.Nonce{
		Address: request.Attributes.Address,
		Message: nonce,
		Expires: time.Now().UTC().Add(30 * time.Minute).Unix(),
	})
	if err != nil {
		logger.WithError(err).Error("failed to insert user nonce")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, resources.NonceResponse{
		Data: resources.Nonce{
			Attributes: resources.NonceAttributes{
				Nonce:   nonce,
				Message: crypto.NonceToMessage(nonce),
			},
		},
	})
}
