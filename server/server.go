package server

import (
	"fmt"
	"log"
	"net/http"
)

func SetupServer(port string) {
	fmt.Printf("Server started at port %s\n", port)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
