package controllers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strconv"
)

type Artist struct {
	Id                    int      `json:"id"`
	Image                 string   `json:"image"`
	Name                  string   `json:"name"`
	Members               []string `json:"members"`
	CreationDate          int      `json:"creationDate"`
	FirstAlbum            string   `json:"firstAlbum"`
	Locations             string   `json:"locations"`
	ConcertDates          string   `json:"concertDates"`
	Relations             string   `json:"relations"`
	Loca                  []string
	Dates                 []string
	RelationsConcertDates map[string][]string
}

func ArtistsManager() ([]Artist, error) {
	var artists []Artist

	linkArtist, er := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if er != nil {
		fmt.Println("Error link artist : ", er)
		return artists, er
	}

	response, _ := io.ReadAll(linkArtist.Body)

	err := json.Unmarshal(response, &artists)
	if err != nil {
		fmt.Println("Error unmarshal artist : ", err)
		return artists, err
	}
	// cela permet de recuperer tout le contenu des locations, des date et relations entre les dates et les concert et de les mettre aux artistes
	locations, erLoc := LocationsManager()
	if erLoc != nil {
		fmt.Println("Error get locations : ", erLoc)
		return artists, erLoc
	}
	dates, erDate := DatesManager()
	if erDate != nil {
		fmt.Println("Error get dates : ", erDate)
		return artists, erDate
	}
	relation, erRelation := RelationsManager()
	if erRelation != nil {
		fmt.Println("Error get relation : ", erRelation)
		return artists, erRelation
	}

	for i := 0; i < len(artists); i++ {
		artists[i].Loca = locations.Index[i].Locations
		artists[i].Dates = dates.Index[i].Dates
		artists[i].RelationsConcertDates = relation.Index[i].DatesLocations

	}
	return artists, nil
}

func GetArtists(w http.ResponseWriter, r *http.Request) {
	artists, er := ArtistsManager()

	if er != nil {
		fmt.Println("Error get artists : ", er)
		Status500(w, r)
		return
	}

	RenderTemplate(w, "artists", artists)
}

var templates = template.Must(template.ParseGlob("client/pages/*.html"))

func RenderTemplate(w http.ResponseWriter, tmpl string, artist interface{}) {
	page := tmpl + ".html"
	err := templates.ExecuteTemplate(w, page, artist)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		tmpl := template.Must(template.ParseFiles("client/pages/500.html"))
		tmpl.Execute(w, struct{ Success bool }{true})
	}
}

func GetArtistInfo(w http.ResponseWriter, r *http.Request) {
	artists, er := ArtistsManager()

	if er != nil {
		fmt.Println("Error get details artists : ", er)
		Status500(w, r)
		return
	}
	// obtenir l'id du artist choisi sur lurl
	artistID := r.URL.Query().Get("id")
	// ID incorrect bad request 400
	id, err := strconv.Atoi(artistID)
	if err != nil {
		Status404(w, r)
		return
	}

	// ID introuvable not found 404
	if id < 1 || id > 52 {
		Status404(w, r)
		return
	}

	artist := artists[id-1]

	// chargement de la page artist
	RenderTemplate(w, "artistDetails", artist)
}
