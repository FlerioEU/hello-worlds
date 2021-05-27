package routes

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/FlerioEU/hello-world/mongo-golang/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r Routes) registerBooks(rg *gin.RouterGroup, db mongo.Database) {
	books := rg.Group("/books")

	books.POST("", postBooks(db))
	books.GET("/:id", getBook(db))
	books.GET("", getBooks(db))
	books.DELETE("/:id", deleteBooks(db))
	books.PUT("/:id", updateBook(db))
}

func postBooks(db mongo.Database) gin.HandlerFunc {
	coll := db.Collection("books")

	return func(c *gin.Context) {
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			log.Printf("Something went wrong while reading body of book: %v\n", err)
			c.String(http.StatusInternalServerError, `{"message": "Something went wrong while posting book. Please try again later`)
			return
		}

		b := models.Book{}
		err = json.Unmarshal(body, &b)
		if err != nil {
			log.Printf("Something went wrong while unmarshalling post book: %v\n", err)
			c.String(http.StatusInternalServerError, `{"message": "Something went wrong while posting book. Please try again later`)
			return
		}

		res, err := coll.InsertOne(context.Background(), b)
		if err != nil {
			log.Printf("An error occured inserting a document into the mongodb: %v\n", err)
			c.String(http.StatusInternalServerError, `{"message": "Something went wrong while posting book. Please try again later`)
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

		var b models.Book
		err = r.Decode(&b)
		if err != nil {
			log.Printf("Something went wrong decoding the book with id '%v': %v\n", id, err)
			c.String(http.StatusInternalServerError, `{"message": "Something went wrong getting the book with id '%v'"}`, id)
			return
		}

		c.JSON(http.StatusOK, b)
	}
}

func getBooks(db mongo.Database) gin.HandlerFunc {
	coll := db.Collection("books")

	return func(c *gin.Context) {
		cursor, err := coll.Find(context.Background(), bson.M{}, nil)
		if err != nil {
			c.String(http.StatusOK, "[]")
			return
		}

		var b []models.Book
		err = cursor.All(context.Background(), &b)
		if err != nil {
			log.Printf("Something went wrong decoding books: %v\n", err)
			c.String(http.StatusInternalServerError, `{"message": "Something went wrong while fetching books. Please try again later`)
			return
		}

		c.JSON(http.StatusOK, b)
	}
}

func updateBook(db mongo.Database) gin.HandlerFunc {
	coll := db.Collection("books")

	return func(c *gin.Context) {
		id := c.Param("id")

		objectId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			c.String(http.StatusBadRequest, `{"message": "Your id is not valid!"}`)
			return
		}

		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			log.Printf("Something went wrong while reading body of book: %v\n", err)
			c.String(http.StatusInternalServerError, `{"message": "Something went wrong while updating book. Please try again later`)
			return
		}

		b := models.Book{}
		err = json.Unmarshal(body, &b)
		if err != nil {
			log.Printf("Something went wrong while unmarshalling update book: %v\n", err)
			c.String(http.StatusInternalServerError, `{"message": "Something went wrong while updating book. Please try again later`)
			return
		}

		r, err := coll.ReplaceOne(
			context.Background(),
			bson.M{"_id": objectId},
			b,
		)
		if err != nil {
			log.Printf("Something went wrong while replacing book: %v\n", err)
			c.String(http.StatusInternalServerError, `{"message": "Something went wrong while updating book. Please try again later`)
			return
		}

		if r.MatchedCount == 0 {
			c.String(http.StatusNotFound, "")
			return
		}

		// update finished - get the newly updated document and provide it to client
		rGet := coll.FindOne(context.Background(), bson.M{"_id": objectId})
		if rGet.Err() != nil {
			log.Printf("Something went wrong while retrieving book after update: %v\n", err)
			c.String(http.StatusInternalServerError, `{"message": "Something went wrong while updating book. Please try again later`)
			return
		}

		err = rGet.Decode(&b)
		if err != nil {
			log.Printf("Something went wrong decoding the book with id '%v': %v\n", id, err)
			c.String(http.StatusInternalServerError, `{"message": "Something went wrong updating the book with id '%v'"}`, id)
			return
		}

		c.JSON(http.StatusOK, b)
	}
}

func deleteBooks(db mongo.Database) gin.HandlerFunc {
	coll := db.Collection("books")

	return func(c *gin.Context) {
		id := c.Param("id")

		objectId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			c.String(http.StatusBadRequest, `{"message": "Your id is not valid!"}`)
			return
		}

		r, err := coll.DeleteOne(context.Background(), bson.M{"_id": objectId})
		if err != nil {
			log.Printf("Something went wrong deleting book with id '%v': %v\n", id, err)
			c.String(http.StatusInternalServerError, `{"message": "Something went wrong while deleting book. Please try again later`)
			return
		}

		if r.DeletedCount == 0 {
			c.String(http.StatusNotFound, "")
			return
		}

		c.String(http.StatusOK, "")
	}
}
