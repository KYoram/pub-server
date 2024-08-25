package apiV1

import (
	consumablesV1 "github.com/KYoram/pub-server/api/v1/consumables"
	"github.com/go-chi/chi/v5"
)

func SetupV1Routes(r chi.Router) {
	r.Group(func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			consumablesV1.SetupConsumablesRoutes(r)
		})
	})
}
