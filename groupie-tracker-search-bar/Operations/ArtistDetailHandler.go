package groupie

import (
	"fmt"
	"net/http"
	"text/template"
)

type ArtistDetailPageData struct {
	Artist    Artist
	Locations []string
	Dates     []string
	Relations map[string][]string
}

func ArtistDetailHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		ServeErrorPage(w, r, 405)
		return
	}

	artistID := r.PathValue("id")

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
	_,_,dates, _,err := Fetch("dates")
	if err != nil {
		ServeErrorPage(w, r, 500)
		return
	}
	_,_,_,relations, err := Fetch("relation")
	if err != nil {
		ServeErrorPage(w, r, 500)
		return
	}

	// Find the selected artist by ID
	found := false
	var selectedArtist Artist
	for _, artist := range artists {
		if artistID == fmt.Sprintf("%d", artist.ID) {
			selectedArtist = artist
			found = true
			break
		}
	}

	// If no artist is found
	if !found {
		ServeErrorPage(w, r, 404)
		return
	}

	// Find the corresponding locations for this artist
	var artistLocations []string
	for _, loc := range locations {
		if loc.ID == selectedArtist.ID {
			artistLocations = loc.Locations
			break
		}
	}

	// Find the corresponding dates for this artist
	var artistDates []string
	for _, date := range dates {
		if date.ID == selectedArtist.ID {
			artistDates = date.Dates
			break
		}
	}

	// Find the corresponding relations (dates and locations) for this artist
	var artistRelations map[string][]string
	for _, relation := range relations {
		if relation.ID == selectedArtist.ID {
			artistRelations = relation.Relations
			break
		}
	}

	// Prepare the data to pass to the template
	data := ArtistDetailPageData{
		Artist:    selectedArtist,
		Locations: artistLocations,
		Dates:     artistDates,
		Relations: artistRelations,
	}

	// Parse and render the artist details template
	tmpl, err := template.ParseFiles("templates/artist_detail.html")
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
