package main

import (
    "github.com/gorilla/mux"
    "os"
    "fmt"
		"net/http"
		"brix-flix-backend/controllers"
		"brix-flix-backend/models"
)

func newRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/movie/new", controllers.CreateMovie).Methods("POST")
	return router
}

func main() {

	models.Init()

	router := newRouter()


	port := os.Getenv("PORT")
	if port == "" {
		port = "10000" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":" + port, router) //Launch the app, visit localhost:10000/api
	if err != nil {
		fmt.Print(err)
	}

}
