package controllers

import (
	"fmt"
	"net/http"
	// "github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// "github.com/jinzhu/gorm"
	"brix-flix-backend/utils"
	// "io/ioutil"
	"brix-flix-backend/models"
	"encoding/json"
)

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	
	movie := &models.Movie{}
	err := json.NewDecoder(r.Body).Decode(movie)
	if err != nil {
		fmt.Println(err)
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}

	resp := movie.Create()
	utils.Respond(w, resp)
}