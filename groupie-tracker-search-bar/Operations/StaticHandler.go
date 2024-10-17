package groupie

import (
	"net/http"
	"os"
)

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	fs := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
	_, err := os.Stat("." + r.URL.Path)
	if err != nil {
		ServeErrorPage(w, r, 404)
		return
	}
	if r.URL.Path == "/static/" {
		ServeErrorPage(w, r, 403)
		return
	}
	fs.ServeHTTP(w, r)
}
