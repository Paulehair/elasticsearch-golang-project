package main

import (
	"github.com/Paulehair/elasticsearch-golang-project/config"
	"github.com/Paulehair/elasticsearch-golang-project/controllers"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Starting server...")
	r := mux.NewRouter()

	log.Println("Connecting to database...")
	config.InitDB()

	r.HandleFunc("/books", controllers.PostBook).Methods("POST")
	r.HandleFunc("/books", controllers.GetBook).Methods("GET")

	log.Println("Server running on localhost:8080!")

	log.Fatal(http.ListenAndServe(":8080", r))
}
