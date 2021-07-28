package controllers

import (
	// "github.com/Paulehair/elasticsearch-golang-project/models"

	// "context"
	// "encoding/json"
	// "log"
	"net/http"
)

func PostBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Post route"))
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Get route"))
}