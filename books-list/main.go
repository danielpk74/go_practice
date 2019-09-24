package main

import (
	"books-list/models"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	// uuid.Must(uuid.NewV4())
)

func init() {
	gotenv.Load()
}

func getDBClient() *mongo.Client {
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URL"))
	client, err := mongo.NewClient(clientOptions)
	logFatal(err)

	err = client.Connect(context.Background())
	logFatal(err)

	return client
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	log.Println("Running GO!")
	log.Fatal(http.ListenAndServe(os.Getenv("APP_PORT"), router))
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Book

	//connn := drivers.MongoDriver{}
	//conn := connn.GetMongoClient()

	conn := getDBClient()
	err := conn.Ping(context.Background(), readpref.Primary())
	logFatal(err)

	collection := conn.Database("books_store").Collection("books")
	cur, err := collection.Find(context.TODO(), bson.D{})
	logFatal(err)

	for cur.Next(context.TODO()) {
		var book models.Book
		err = cur.Decode(&book)
		logFatal(err)

		books = append(books, book)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
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

func getBook(w http.ResponseWriter, r *http.Request) {
	var (
		book      models.Book
		params    = mux.Vars(r)
		bookId, _ = strconv.Atoi(params["id"])
	)

	conn := getDBClient()
	err := conn.Ping(context.Background(), readpref.Primary())
	logFatal(err)

	collection := conn.Database("books_store").Collection("books")
	filter := bson.D{{"id", bookId}}
	cur := collection.FindOne(context.TODO(), filter)
	err = cur.Decode(&book)

	// err = collection.FindOne(context.TODO(), filter).Decode(&book)
	// if book == nil {
	// 	log.Println("book not found.")
	// }

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(book)
}

func addBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)

	conn := getDBClient()
	err := conn.Ping(context.Background(), readpref.Primary())
	logFatal(err)

	collection := conn.Database("books_store").Collection("books")
	insertResult, err := collection.InsertOne(context.TODO(), book)
	logFatal(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	json.NewEncoder(w).Encode(&insertResult)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	var (
		params    = mux.Vars(r)
		bookId, _ = strconv.Atoi(params["id"])
		book      models.Book
	)

	json.NewDecoder(r.Body).Decode(&book)

	conn := getDBClient()
	err := conn.Ping(context.Background(), readpref.Primary())
	logFatal(err)

	collection := conn.Database("books_store").Collection("books")
	filter := bson.D{{"id", bookId}}
	updated := bson.D{{Key: "$set", Value: book}}
	updateResult, err := collection.UpdateOne(context.TODO(), filter, updated)
	logFatal(err)

	json.NewEncoder(w).Encode(updateResult)
}

func removeBook(w http.ResponseWriter, r *http.Request) {
	var (
		params    = mux.Vars(r)
		bookId, _ = strconv.Atoi(params["id"])
	)

	conn := getDBClient()
	err := conn.Ping(context.Background(), readpref.Primary())
	logFatal(err)

	collection := conn.Database("books_store").Collection("books")
	filter := bson.D{{"id", bookId}}
	resultDeletion, err := collection.DeleteOne(context.TODO(), filter)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(&resultDeletion)
}

// books = append(books, Book{ID: 1, Title: "The Yisus Book", Author: "Yisus", Year: "2015"},
// 		Book{ID: 2, Title: "The Yisus Book II", Author: "Yisus", Year: "2016"},
// 		Book{ID: 3, Title: "The Yisus Book III", Author: "Yisus", Year: "2017"},
// 		Book{ID: 4, Title: "The Yisus Book IV", Author: "Yisus", Year: "2018"})

//log.Println(reflect.TypeOf(i))
//log.Println(params)
