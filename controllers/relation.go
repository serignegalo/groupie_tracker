package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Relation struct {
	Index []struct {
		Id             int                 `json:"id"`
		DatesLocations map[string][]string `json:"dateslocations"`
	} `json:"index"`
}

func RelationsManager() (Relation, error) {
	var relation Relation

	link, er := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if er != nil {
		fmt.Println("Error link relation : ", er)
		return relation, er
	}

	response, _ := io.ReadAll(link.Body)

	err := json.Unmarshal(response, &relation)
	if err != nil {
		fmt.Println("Error unmarshal relation : ", err)
		return relation, err
	}

	return relation, nil
}

func GetRelations(w http.ResponseWriter, r *http.Request) {
	relation, er := RelationsManager()
	if er != nil {
		fmt.Println("Error get relation : ", er)
		Status500(w, r)
		return
	}
	//chargement de la page relation
	RenderTemplate(w, "relation", relation.Index)
}
