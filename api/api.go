package api

import (
	"net/http"
	"regexp"

	apiV1 "github.com/KYoram/pub-server/v1/api"
)

func HandleApi(response http.ResponseWriter, request *http.Request) {
	regex, err := regexp.Compile(`\/api\/v1(\/?(.+))?`)
	if err != nil {
		http.Error(response, "500 issue", http.StatusInternalServerError)
	}

	if regex.Match([]byte(request.URL.Path)) {
		apiV1.HandleApiV1(response, request)
		return
	}

	http.Redirect(response, request, "/api/v1", http.StatusPermanentRedirect)
}
