package main

import (
	"net/http"

	"github.com/KYoram/pub-server/api"
	"github.com/KYoram/pub-server/file"
	"github.com/KYoram/pub-server/server"
)

func main() {
	http.Handle("/", file.LoadStatic())
	http.HandleFunc("/api/", api.HandleApi)
	server.SetupServer("8080")
}
