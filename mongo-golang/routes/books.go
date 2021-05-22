package routes

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type book struct {
	Id     string `bson:"_id,omitempty"`
	Title  string `bson:"title,omitempty"`
	Author string `bson:"author,omitempty"`
}

func (r Routes) registerBooks(rg *gin.RouterGroup, db mongo.Database) {
	books := rg.Group("/books")

	books.POST("", postBooks(db))
	books.GET("/:id", getBook(db))
}

func postBooks(db mongo.Database) gin.HandlerFunc {
	coll := db.Collection("Books")

	return func(c *gin.Context) {
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			log.Println("Could not read the body of request!")
			return
		}

		b := book{}
		err = json.Unmarshal(body, &b)
		if err != nil {
			log.Println("Could not unmarshal the body of request!")
			return
		}

		res, err := coll.InsertOne(context.Background(), b)
		if err != nil {
			log.Printf("An error occured inserting a document into the mongodb: %v\n", err)
			return
		}

		r := coll.FindOne(context.Background(), bson.M{"_id": res.InsertedID})
		r.Decode(&b)
		c.JSON(http.StatusCreated, b)
	}
}

func getBook(db mongo.Database) gin.HandlerFunc {
	coll := db.Collection("books")

	return func(c *gin.Context) {
		id := c.Param("id")

		objectId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			c.String(http.StatusBadRequest, `{"message": "Your id is not valid!"}`)
			return
		}

		r := coll.FindOne(context.Background(), bson.M{"_id": objectId})
		if r.Err() != nil {
			c.String(http.StatusNotFound, "")
			return
		}

		var b book
		err = r.Decode(&b)
		if err != nil {
			log.Printf("Something went wrong decoding the book with id '%v': %v\n", id, err)
			c.String(http.StatusInternalServerError, `{"message": "Something went wrong getting the book with id '%v'"}`, id)
			return
		}

		c.JSON(http.StatusOK, b)
	}
}
