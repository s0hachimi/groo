package main

import (
	"fmt"
	"log"
	"net/http"

	x "groupie/Operations"
)

func main() {
	http.HandleFunc("/", x.HomeHandler)               
	http.HandleFunc("/Artist/{id}", x.ArtistDetailHandler)
	http.HandleFunc("/searchBar", x.SearchBar)
	http.HandleFunc("/static/", x.StaticHandler)
	fmt.Println("http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("error starting server : ", err)
	}
}
