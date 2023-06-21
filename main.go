package main

import (
	"fmt"
	"net/http"
	"groupie-tracker/controllers"
)

func main() {

	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/artist", controllers.GetArtists)
	http.HandleFunc("/artistDetails", controllers.GetArtistInfo)
	http.HandleFunc("/relation", controllers.GetRelations)
	http.HandleFunc("/location", controllers.GetLocations)
	http.HandleFunc("/date", controllers.GetDates)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("client/css/"))))
	fmt.Println("Server starting on port 8080: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
