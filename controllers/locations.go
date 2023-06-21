package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Location struct {
	Index []struct {
		Id        int      `json:"id"`
		Locations []string `json:"locations"`
		Dates     string   `json:"dates"`
	} `json:"index"`
}

func LocationsManager() (Location, error) {
	var locations Location

	linkLocation, er := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if er != nil {
		fmt.Println("Error link location: ", er)
		return locations, er
	}

	response, _ := io.ReadAll(linkLocation.Body)

	err := json.Unmarshal(response, &locations)
	if err != nil {
		fmt.Println("Error unmarshal locations : ", err)
		return locations, err
	}

	return locations, nil
}

func GetLocations(w http.ResponseWriter, r *http.Request) {
	location, er := LocationsManager()
	if er != nil {
		fmt.Println("Error get location : ", er)
		Status500(w, r)
		return
	}
	//chargement de la page relation
	RenderTemplate(w, "location", location.Index)
}
