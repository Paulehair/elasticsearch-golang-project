package controllers

import (
	"github.com/Paulehair/elasticsearch-golang-project/config"
	"github.com/Paulehair/elasticsearch-golang-project/models"

	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func PostBook(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)

	if err != nil {
		log.Printf("Err decode body: %s", err.Error())
		json.NewEncoder(w).Encode("Error decoding body")
		return
	}

	res, err := config.CreateBook(ctx, &book)

	if err != nil {
		log.Printf("Error create book: %s", err.Error())
		return
	}

	json.NewEncoder(w).Encode(fmt.Sprintf("Sucessful insertion: %v", res))
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var queryKey = r.URL.Query().Get("key")
	var queryValue = r.URL.Query().Get("value")

	if len(queryKey) > 0 && len(queryValue) > 0 {
		books, err := config.GetBooks(ctx, queryKey, queryValue)

		if err != nil {
			json.NewEncoder(w).Encode(err.Error())
			log.Printf("Error reading book: %s", err.Error())
			return
		}

		json.NewEncoder(w).Encode(books)
	} else {
		json.NewEncoder(w).Encode(`{ "error": "Missing key or value parameters. Please give :(" }`)
	}

}
