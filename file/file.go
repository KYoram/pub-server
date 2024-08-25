package file

import "net/http"

func ServeStatic() http.Handler {
	var fileServer = http.FileServer(http.Dir("./bin/static"))

	return fileServer
}
