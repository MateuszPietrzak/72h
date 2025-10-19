package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"

	"github.com/MateuszPietrzak/72h/templates/pages"
	"github.com/a-h/templ"
)

var dev = os.Getenv("ENV") != "production"

func disableCacheInDevMode(next http.Handler) http.Handler {
	if !dev {
		return next
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		next.ServeHTTP(w, r)
	})
}

func main() {

	fs := http.FileServer(http.Dir("./static"))

	http.Handle("/static/",
		disableCacheInDevMode(
			http.StripPrefix("/static/", fs)))

	http.Handle("/", templ.Handler(pages.Home()))
	http.Handle("/tech_campy", templ.Handler(pages.TechCamps()))

	h := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Pong!")
	}
	http.HandleFunc("/ping", h)

	socketPath := os.Getenv("UNIX_SOCKET")
	useSocket := socketPath != ""
	if useSocket {
		os.Remove(socketPath)

		listener, err := net.Listen("unix", socketPath)
		if err != nil {
			panic(err)
		}
		defer listener.Close()

		os.Chmod(socketPath, 0660)

		fmt.Println("Listening on Unix socket", socketPath)
		http.Serve(listener, nil)
	} else {
		port := os.Getenv("HTTP_PORT")
		if port == "" {
			port = "8080"
		}
		fmt.Println("Listening on port", port)
		http.ListenAndServe(":"+port, nil)
	}
}
