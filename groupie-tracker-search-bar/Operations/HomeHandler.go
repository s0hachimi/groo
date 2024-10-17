package groupie

import (
	"net/http"
	"text/template"
)

type ArtistsPageData struct {
	Artists   []Artist
	Locations []string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		ServeErrorPage(w, r, 405)
		return
	}

	if r.URL.Path != "/" {
		ServeErrorPage(w, r, 404)
		return
	}

	artists,_,_,_, err := Fetch("artist")
	if err != nil {
		ServeErrorPage(w, r, 500)
		return
	}

	_,locations,_,_, err := Fetch("location")
	if err != nil {
		ServeErrorPage(w, r, 500)
		return
	}
	seen := make(map[string]bool)
	var locati []string
	for _,loca := range locations{
		for _, r := range loca.Locations {
			if !seen[r] {
				locati = append(locati, r)
				seen[r] = true
			}
		}
	}
	// Prepare the data with only artist images for the home page
	data := ArtistsPageData{
		Artists:   artists,
		Locations: locati,
	}

	// Parse and execute the template for the homepage (artist images)
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		ServeErrorPage(w, r, 500)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		ServeErrorPage(w, r, 500)
		return
	}
}
