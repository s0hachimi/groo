package groupie

import (
	"encoding/json"
	"io"
	"net/http"
)

type Artist struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}

type Date struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Relation struct {
	ID        int                 `json:"id"`
	Relations map[string][]string `json:"datesLocations"`
}
var (
	LocationURL = "https://groupietrackers.herokuapp.com/api/locations"
	ArtistURL = "https://groupietrackers.herokuapp.com/api/artists"
	DatesURL =  "https://groupietrackers.herokuapp.com/api/dates"
	RelationURL = "https://groupietrackers.herokuapp.com/api/relation"
)

var (
	artists []Artist
	Loca    struct {
		Index []Location `json:"index"`
	}
	Dates struct {
		Index []Date `json:"index"`
	}
	Rela struct {
		Index []Relation `json:"index"`
	}
)
func Fetch(pattern string) ([]Artist, []Location, []Date, []Relation, error) {
	var artist []Artist
	var location []Location
	var date []Date
	var relation []Relation
	url := ""
	if pattern == "artist" {
		url = ArtistURL
	} else if pattern == "location" {
		url = LocationURL
	} else if pattern == "dates" {
		url = DatesURL
	} else if pattern == "relation" {
		url = RelationURL
	}
	resp, err := http.Get(url)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	if pattern == "artist" {
		err = json.Unmarshal(body, &artists)
		if err != nil {
			return nil, nil, nil, nil, err
		}
		artist = artists
	} else if pattern == "location" {
		err = json.Unmarshal(body, &Loca)
		if err != nil {
			return nil, nil, nil, nil, err
		}
		location = Loca.Index
	} else if pattern == "dates" {
		err = json.Unmarshal(body, &Dates)
		if err != nil {
			return nil, nil, nil, nil, err
		}
		date = Dates.Index
	} else if pattern == "relation" {
		err = json.Unmarshal(body, &Rela)
		if err != nil {
			return nil, nil, nil, nil, err
		}
		relation = Rela.Index
	}
	return artist, location, date, relation, nil
}