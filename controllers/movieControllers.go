package controllers

import (
	"fmt"
	"net/http"
	// "github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jinzhu/gorm"
	"brix-flix-backend/utils"
	// "io/ioutil"
	"brix-flix-backend/models"
	"encoding/json"
)

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	dbUri := utils.DbString()
	db, err := gorm.Open("postgres", dbUri)
	if err != nil {
		panic("Could not connect.")
	}
	defer db.Close()

	defer r.Body.Close()
	var movies models.Movies
	details := json.NewDecoder(r.Body).Decode(&movies);
	fmt.Println(details)
	// err != nil {
		// respondWithError(w, http.StatusBadRequest, "Invalid request payload")
	// 	return
	// }
	// body, err := json.Marshal(r.Body)
	// data, _ := ioutil.ReadAll(r.Body)
	// fmt.Println(string(data))

	// vars := mux.Vars(r)
	// fmt.Println(vars)



	// movies := models.Movies{}
	// details := json.NewDecoder(r.Body).Decode(&movies)

	// details, err := json.Marshal(movies)
	// jsonValue, _ := json.Marshal(movies)
	// details := json.NewDecoder(r.Body).Decode(&models.Movies{})

	// fmt.Print(details)
	// db.Create(&models.Movies{Name: name, Year: year, Description: description})
	
	fmt.Fprintf(w, "New movie added!")
}