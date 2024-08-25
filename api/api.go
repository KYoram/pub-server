package api

import (
	apiV1 "github.com/KYoram/pub-server/api/v1"
	"github.com/go-chi/chi/v5"
)

func SetupApiRoutes(r chi.Router) {
	r.Group(func(r chi.Router) {
		r.Route("/api", func(r chi.Router) {
			apiV1.SetupV1Routes(r)
		})
	})
}
