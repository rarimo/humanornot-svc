package requests

import (
	"encoding/json"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"
	"gitlab.com/rarimo/identity/kyc-service/internal/crypto"
	"gitlab.com/rarimo/identity/kyc-service/resources"
	"net/http"
	"strings"
)

type NonceRequest struct {
	Data resources.NonceRequest `json:"data"`
}

func NewNonceRequest(r *http.Request) (resources.NonceRequest, error) {
	var request NonceRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return resources.NonceRequest{}, errors.Wrap(err, "failed to unmarshal")
	}

	request.Data.Attributes.Address = strings.ToLower(request.Data.Attributes.Address)

	return request.Data, request.validate()
}

func (r *NonceRequest) validate() error {
	return validation.Errors{
		"/data/type": validation.Validate(
			r.Data.Type,
			validation.Required,
			validation.In(resources.NONCE_REQUEST),
		),
		"/data/attributes/address": validation.Validate(
			r.Data.Attributes.Address,
			validation.Required,
			validation.Match(crypto.AddressRegexp),
		),
	}.Filter()
}
