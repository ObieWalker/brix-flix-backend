package main

import (
    "github.com/gorilla/mux"
    "os"
    "fmt"
		"net/http"
		"brix-flix-backend/controllers"
		"brix-flix-backend/models"

    // importing any module from the root should
    // be easier to do e.g: models "brix_flix_server/models"
)

// func homePage(w http.ResponseWriter, r *http.Request){
//     fmt.Fprintf(w, "Welcome to the HomePage!")
//     fmt.Println("Endpoint Hit: homePage")
// }

// func handleRequests() {
//     http.HandleFunc("/", homePage)
//     log.Fatal(http.ListenAndServe(":10000", nil))
// }

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
