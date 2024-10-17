package groupie

import (
	"fmt"
	"net/http"
	"strings"
	"text/template"
)

type Art struct {
	Nart []Artist
}

func SearchBar(w http.ResponseWriter, r *http.Request) {
	search := strings.ToLower(r.FormValue("search"))
	if len(search) == 0 {
		ServeErrorPage(w, r, 400)
		return
	}

	artist, _, _, _, err := Fetch("artist")
	if err != nil {
		ServeErrorPage(w, r, 500)
		return
	}

	_, locations, _, _, err := Fetch("location")
	if err != nil {
		ServeErrorPage(w, r, 500)
		return
	}

	var NewArt []Artist

	found := map[int]bool{}
	for _, artist := range artist {
		if strings.Contains(strings.ToLower(artist.Name), search) {
			NewArt = append(NewArt, artist)
			found[artist.ID] = true
		}
		for _, Members := range artist.Members {
			if strings.Contains(strings.ToLower(Members), search) {
				if !found[artist.ID] {
					NewArt = append(NewArt, artist)
					found[artist.ID] = true
				}
			}
		}
		if strings.Contains(fmt.Sprintf("%v", artist.CreationDate), search) {
			if !found[artist.ID] {
				NewArt = append(NewArt, artist)
				found[artist.ID] = true
			}
		}
		if strings.Contains(strings.ToLower(artist.FirstAlbum), search) {
			if !found[artist.ID] {
				NewArt = append(NewArt, artist)
				found[artist.ID] = true
			}
		}
		for _, location := range locations {
			for _, loc := range location.Locations {
				if strings.Contains(strings.ToLower(loc), search) {
					if artist.ID == location.ID && !found[artist.ID] {
						NewArt = append(NewArt, artist)
						found[artist.ID] = true
					}
				}
			}
		}
	}

	if len(NewArt) == 0 {
		ServeErrorPage(w, r, 404)
		return
	}

	tmpl, err := template.ParseFiles("templates/search-bar.html")
	if err != nil {
		ServeErrorPage(w, r, 500)
		return
	}

	data := Art{
		Nart: NewArt,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		ServeErrorPage(w, r, 500)
		return
	}
}
