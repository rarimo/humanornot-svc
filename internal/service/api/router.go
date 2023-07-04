package api

import (
	"fmt"
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"

	"gitlab.com/rarimo/identity/kyc-service/internal/service/api/handlers"
	"gitlab.com/rarimo/identity/kyc-service/internal/service/api/requests"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
			handlers.CtxKYCService(s.kycService),
			handlers.CtxMasterQueryer(s.masterQ),
		),
	)

	r.Route("/integrations/kyc-service", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/public", func(r chi.Router) {
				r.Post(fmt.Sprintf("/verify/{%s}", requests.IdentityProviderPathParam), handlers.Verify)
				r.Post("/nonce", handlers.GetNonce)
				r.Get("/status/{verification-id}", handlers.GetVerifyStatus)
			})
		})
	})

	return r
}
