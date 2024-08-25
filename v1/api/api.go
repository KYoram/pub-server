package apiV1

import (
	"fmt"
	"net/http"
)

func HandleApiV1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "success")
}
