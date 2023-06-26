package api

import (
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"

	"gitlab.com/rarimo/identity/kyc-service/internal/service/api/handlers"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
		),
	)
	r.Route("/integrations/kyc-service", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {

		})
	})

	return r
}
