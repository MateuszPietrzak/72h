package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"

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

func staticScript(path string) string {
	srcpath := filepath.Join("static/scripts", path, "script.html")
	dat, err := os.ReadFile(srcpath)
	if err != nil {
		log.Fatalf("Failed to read static script file: %v", err)
	}
	return string(dat)
}

func main() {

	fs := http.FileServer(http.Dir("./static"))

	http.Handle("/static/",
		disableCacheInDevMode(
			http.StripPrefix("/static/", fs)))

	http.Handle("/", templ.Handler(pages.Home()))
	http.Handle("/scenariusze", templ.Handler(pages.ScriptsHome()))
	http.Handle("/tech_campy", templ.Handler(pages.TechCamps()))
	http.Handle("/aplikacja", templ.Handler(pages.App()))
	http.Handle("/edukacja", templ.Handler(pages.Education()))
	http.Handle("/o_nas", templ.Handler(pages.AboutUs()))

	http.Handle("/scenariusze/1/1", templ.Handler(pages.Script(staticScript("Moduł 1 - Zespół w akcji/1. Budowanie zespołu"))))
	http.Handle("/scenariusze/1/2", templ.Handler(pages.Script(staticScript("Moduł 1 - Zespół w akcji/2. Komunikacja w zespole"))))
	http.Handle("/scenariusze/1/3", templ.Handler(pages.Script(staticScript("Moduł 1 - Zespół w akcji/3. Motywacja zespołu"))))
	http.Handle("/scenariusze/1/4", templ.Handler(pages.Script(staticScript("Moduł 1 - Zespół w akcji/4. Planowanie działania w zespole"))))
	http.Handle("/scenariusze/1/5", templ.Handler(pages.Script(staticScript("Moduł 1 - Zespół w akcji/5. Wybór lidera"))))

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
