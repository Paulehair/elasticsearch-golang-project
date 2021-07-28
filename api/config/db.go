package config

import (
	"github.com/Paulehair/elasticsearch-golang-project/models"
	"github.com/olivere/elastic/v7"

	"context"
	"log"
	"time"
	"errors"
)

// ES client setup

var EsClient = *elastic.Client

func InitDB() {
	ctx := context.Background()

	StartESClient()

	for {
		exists, err := EsClient.IndexExists("books").Do(ctx)

		if err != nil {
			log.Println("Error:", err)
			continue
		}

		if !exists {
			_, err := EsClient.CreateIndex("books").BodyString(models.BookMapping).Do(ctx)
			if err != nil {
				log.Println("Error:", err)
			}
			log.Println("ES connected.")
			break
		}

		log.Println("ES connected.")
		break
	}
}

func StartESClient() {
	for {
		log.Println("Connecting to ES")

		client, err := elastic.NewClient(elastic.SetURL("http://elasticsearch:9200"), elastic.SetSniff(false), elastic.SetHealthcheck(false))

		if err != nil {
			log.Println("Failed to connect, retrying in 5 seconds...")
			time.sleep(5 * time.Second)
			continue
		}

		EsClient = client
		log.Println("ES connected.")
		return
	}
}

// ES methods

func prepareForES(value interface{}) (string, error) {
	b, err := json.Marshal(value)
	if err != nil {
		return "", errors.New("error marshal book")
	}

	return string(b), nil
}

func CreateBook(ctx context.Context, book *models.Book) (string, error) {
	s, err := prepareForES(book)

	if err != nil {
		return "", errors.New("Error creating item")
	}

	res, err := EsClient.Index().Index("books").BodyJson(s).Do(ctx)
	if err != nil {
		return "", errors.New("error insert book")
	}
	return res.Result, nil
}

func SearchBooks(ctx context.Context, key, value string) ([]models.Book, error) {
	searchSource := elastic.NewSearchSource()
	searchSource.Query(elastic.NewMatchQuery(key, value))

	searchRes, err := EsClient.Search().Index("books").SearchSource(searchSource).Do(ctx)
	if err != nil {
		log.Println("Error searching book")
		return nil, err
	}

	var books []models.Book
	for _, res := range searchRes.Hits.Hits {
		var book models.Book
		err := json.Unmarshal(res.Source, &book)
		if err != nil {
			return nil, errors.New("cannot decode book")
		}

		books = append(books, book)
	}

	return books, nil
}