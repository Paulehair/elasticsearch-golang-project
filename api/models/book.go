package models

type Book struct {
	ID        		string `json:"id"`
	Title     		string `json:"title"`
	Abstract  		string `json:"abstract"`
	Author    		string `json:"author"`
	Category  		string `json:"category"`
	Isbn      		string `json:"isbn"`
	PublicationDate string `json:"publicationDate"`
}

const BookMapping = `{
	"settings": {
		"number_of_shards": 1,
		"number_of_replicas": 1
	},
	"mappings": {
		"properties": {
			"name": {
				"type": "text"
			},
			"title": {
				"type": "text"
			},
			"author": {
				"type": "text"
			},
			"abstract": {
				"type": "text"
			},
			"category": {
				"type": "text"
			},
			"isbn": {
				"type": "text"
			}
		}
	}
}`