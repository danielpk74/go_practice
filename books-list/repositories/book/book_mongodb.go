package bookrepository

import (
	"books-list/models"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type BookRepository struct{}

// Description comment
func (br BookRepository) GetBooks() ([]models.Book, error) {
	conn := getDBClient()
	err := conn.Ping(context.Background(), readpref.Primary())
	log.Fatal(err)

	collection := conn.Database("books_store").Collection("books")
	cur, err := collection.Find(context.TODO(), bson.D{})
	log.Fatal(err)

	for cur.Next(context.TODO()) {
		var book models.Book
		err = cur.Decode(&book)
		log.Fatal(err)

		books = append(books, book)
	}
}
