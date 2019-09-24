package controllers

import (
	"books-list/models"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Controller struct{}

var books []models.Book

// Returns a JSON Object with all books from database
func (c Controller) GetBooks(dbClient *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var books []models.Book

		err := dbClient.Ping(context.Background(), readpref.Primary())
		log.Fatal(err)

		collection := dbClient.Database("books_store").Collection("books")
		cur, err := collection.Find(context.TODO(), bson.D{})
		log.Fatal(err)

		for cur.Next(context.TODO()) {
			var book models.Book
			err = cur.Decode(&book)
			log.Fatal(err)

			books = append(books, book)
		}

		json.NewEncoder(w).Encode(books)

		// // Connect the mongo client to the MongoDB server
		// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		// err := mongoClient.Connect(ctx)
		// logFatal(err)

		// collection := mongoClient.Database("books_store").Collection("books")
		// cur, err := collection.Find(ctx, bson.D{})
		// logFatal(err)

		// defer cur.Close(ctx)
		// for cur.Next(ctx) {
		// 	var result bson.RawValue
		// 	err := cur.Decode(&result)
		// 	log.Println(&result)
		// 	logFatal(err)
		// }

	}
}
