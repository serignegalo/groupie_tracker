package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Date struct {
	Index []struct {
		Id    int      `json:"id"`
		Dates []string `json:"dates"`
	} `json:"index"`
}

func DatesManager() (Date,error){
	var dates Date

	link, er := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	if er != nil {
		fmt.Println("Error to read  dates: ", er)
		return dates,er
	}

	response, _ := io.ReadAll(link.Body)


	err := json.Unmarshal(response, &dates)
	if err != nil {
		fmt.Println("Error unmarshal : ", err)
		return dates, err
	}

	//fmt.Println(dates)

	return dates, nil
}

func GetDates(w http.ResponseWriter, r *http.Request) {
	dates, er := DatesManager()
	if er != nil {
		fmt.Println("Error get dates : ", er)
		Status500(w, r)
		return
	}
	//chargement de la page relation
	RenderTemplate(w, "date", dates.Index)
}

