package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// https://dev.to/johanlejdung/a-mini-guide-build-a-rest-api-as-a-go-microservice-together-with-mysql-27m2

var collection *mongo.Collection

type book struct {
	Title  string
	Author string
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	uri := os.Getenv("MONGO_URL")

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("An error occured while connecting to the mongodb: %v", err)
	}

	collection = client.Database("hw-golang").Collection("books")

	http.HandleFunc("/books", postBook)
	http.HandleFunc("/books/", getBook)

	log.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

func postBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Write([]byte("Method not allowed!"))
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Could not read the body of request!")
	}

	b := book{}
	err = json.Unmarshal(body, &b)
	if err != nil {
		log.Println("Could not unmarshal the body of request!")
	}

	res, err := collection.InsertOne(context.Background(), b)
	if err != nil {
		log.Fatalf("An error occured inserting a document into the mongodb: %v\n", err)
	}

	log.Printf("An object with the id '%v' has been created!\n", res.InsertedID)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Write([]byte("Method not allowed!"))
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/books/")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("%v is not a valid ObjectId\n", id)
	}

	result := collection.FindOne(context.Background(), bson.M{"_id": objectId})
	if result.Err() != nil {
		log.Printf("An error occured while trying to fetch an object with id %v: %v\n", id, result.Err())
	}

	var b book
	err = result.Decode(&b)
	if err != nil {
		log.Printf("An error occured while trying to decode an object with id %v: %v\n", id, err)
	}

	json, err := json.Marshal(b)
	if err != nil {
		log.Printf("An error occured while trying to marshal to json %v\n", err)
	}

	w.Write(json)
}
