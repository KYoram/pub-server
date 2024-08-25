package consumablesV1

import "net/http"

func get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get"))
}

func update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("update"))
}

func delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete"))
}
