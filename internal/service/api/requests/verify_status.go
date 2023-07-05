package requests

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
)

type VerifyStatusRequest struct {
	VerifyID uuid.UUID
}

type verifyStatusRequest struct {
	VerifyID string
}

func NewVerifyStatusRequest(r *http.Request) (*VerifyStatusRequest, error) {
	requestRaw := verifyStatusRequest{
		VerifyID: chi.URLParam(r, VerifyIDPathParam),
	}

	if err := requestRaw.validate(); err != nil {
		return nil, err
	}

	return requestRaw.parse(), nil
}

func (req *verifyStatusRequest) validate() error {
	return validation.Errors{
		fmt.Sprintf("path/%s", VerifyIDPathParam): validation.Validate(
			req.VerifyID, validation.Required, validation.By(MustBeValidUUID),
		),
	}.Filter()
}

func (req *verifyStatusRequest) parse() *VerifyStatusRequest {
	verifyID, _ := uuid.Parse(req.VerifyID)

	return &VerifyStatusRequest{
		VerifyID: verifyID,
	}
}

func MustBeValidUUID(src interface{}) error {
	uuidRaw, ok := src.(string)
	if !ok {
		return errors.New("it is not a string")
	}

	_, err := uuid.Parse(uuidRaw)
	if err != nil {
		return errors.New("it is not a valid uuid")
	}

	return nil
}
