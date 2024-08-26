package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/LeRoiDesPoissons/pub-server/api"
	"github.com/LeRoiDesPoissons/pub-server/file"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

var AdminCode int = rand.New(rand.NewSource(time.Now().UnixNano())).Int()
var ClientCode int = rand.New(rand.NewSource(time.Now().UnixNano())).Int()

func main() {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Use(render.SetContentType(render.ContentTypeJSON))

	router.Use(middleware.Timeout(60 * time.Second))

	api.SetupApiRoutes(router)

	router.Handle("/*", file.ServeStatic())

	server := httptest.NewServer(router)
	port := strings.Split(server.URL, ":")[2]

	fmt.Println("Server starting at " + server.URL)
	fmt.Printf("Use code %d to log in\n", AdminCode)

	openURL(server.URL)

	err := http.ListenAndServe(":"+port, router)

	if err != nil {
		panic(err)
	}
}

func openURL(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start", url}
	case "darwin":
		cmd = "open"
		args = []string{url}
	default:
		if isWSL() {
			cmd = "cmd.exe"
			args = []string{"/c", "start", url}
		} else {
			cmd = "xdg-open"
			args = []string{url}
		}
	}
	if len(args) > 1 {
		args = append(args[:1], append([]string{""}, args[1:]...)...)
	}
	return exec.Command(cmd, args...).Start()
}

func isWSL() bool {
	releaseData, err := exec.Command("uname", "-r").Output()
	if err != nil {
		return false
	}
	return strings.Contains(strings.ToLower(string(releaseData)), "microsoft")
}
