package file

import "net/http"

func LoadStatic() http.Handler {
	var fileServer = http.FileServer(http.Dir("./bin/static"))

	return fileServer
}
