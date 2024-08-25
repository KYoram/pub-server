package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/KYoram/pub-server/api"
	"github.com/KYoram/pub-server/file"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/docgen"
	"github.com/go-chi/render"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Use(render.SetContentType(render.ContentTypeJSON))

	router.Use(middleware.Timeout(60 * time.Second))

	router.Handle("/", file.ServeStatic())

	api.SetupApiRoutes(router)

	docgen.PrintRoutes(router)

	fmt.Println("Listening on port 8080")

	log.Fatal(http.ListenAndServe(":8080", router))
}
