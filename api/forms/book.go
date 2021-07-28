package forms

import (
	"github.com/Paulehair/elasticsearch-golang-project/models"
)

type BookForm struct {
	Isbn        string `json:"isbn"`
	Title       string `json:"title"`
	Author      string `json:"Author"`
	Description string `json:"description"`
	AmazonURL   string `json:"amazon_url"`
}

func (b BookForm) PrepareBook(book models.Book) models.Book {
	book_data := models.Book{
		Isbn:        b.Isbn,
		Title:       b.Title,
		Author:      b.Author,
		Description: b.Description,
		AmazonURL:   b.AmazonURL,
	}
	return book_data
}
