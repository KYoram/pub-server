package consumablesV1

import (
	"github.com/go-chi/chi/v5"
)

func SetupConsumablesRoutes(r chi.Router) {
	r.Group(func(r chi.Router) {
		setupListRoutes(r)
		setupByIdRoutes(r)
	})
}

func setupListRoutes(r chi.Router) {
	r.Route("/consumables", func(r chi.Router) {
		r.Get("/", list)
		r.Post("/", create)
		r.Get("/search", search)
	})
}

func setupByIdRoutes(r chi.Router) {
	r.Route("/consumable", func(r chi.Router) {
		r.Route("/{consumableID}", func(consumableByIdRouter chi.Router) {
			consumableByIdRouter.Get("/", get)
			consumableByIdRouter.Put("/", update)
			consumableByIdRouter.Delete("/", delete)
		})
	})
}
