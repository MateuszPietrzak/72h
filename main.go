package main

import (
	"fmt"
	"net/http"

	"github.com/MateuszPietrzak/72h/templates/pages"
	"github.com/a-h/templ"
)

var dev = true

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

	home := pages.Home()
	tech_camps := pages.TechCamps()
	scripts_home := pages.ScriptsHome()

	http.Handle("/", templ.Handler(home))
	http.Handle("/scenariusze", templ.Handler(scripts_home))
	http.Handle("/tech_campy", templ.Handler(tech_camps))

	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}
