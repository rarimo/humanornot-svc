package requests

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	core "github.com/iden3/go-iden3-core/v2"
)

type GetProviderByIdentityIdRequest struct {
	IdentityID core.ID
}

type getProviderByIdentityIdRequest struct {
	IdentityId string
}

func NewGetProviderByIdentityIdRequest(r *http.Request) (*GetProviderByIdentityIdRequest, error) {
	/* because in generated resources response type == request type */
	requestRaw := getProviderByIdentityIdRequest{
		IdentityId: chi.URLParam(r, GetProviderByIdentityIdPathParam),
	}

	if err := requestRaw.validate(); err != nil {
		return nil, err
	}

	return requestRaw.parse(), nil
}

func (req *getProviderByIdentityIdRequest) validate() error {
	return validation.Errors{
		fmt.Sprintf("path/%s", GetProviderByIdentityIdPathParam): validation.Validate(
			req.IdentityId, validation.Required, validation.By(MustBeValidID),
		),
	}.Filter()
}

func (req *getProviderByIdentityIdRequest) parse() *GetProviderByIdentityIdRequest {
	identityID, _ := core.IDFromString(req.IdentityId)

	return &GetProviderByIdentityIdRequest{
		IdentityID: identityID,
	}
}
