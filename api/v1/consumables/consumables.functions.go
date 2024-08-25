package consumablesV1

import (
	"net/http"
)

func list(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("list"))
}

func create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create"))
}

func search(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("search"))
}
