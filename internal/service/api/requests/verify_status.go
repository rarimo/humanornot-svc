package requests

import (
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"net/http"
)

func NewVerifyStatusRequest(r *http.Request) (uuid.UUID, error) {
	id := chi.URLParam(r, "verification-id")

	parsed, err := uuid.Parse(id)
	if err != nil {
		return parsed, errors.Wrap(err, "failed to parse parsed")
	}

	return parsed, nil
}
