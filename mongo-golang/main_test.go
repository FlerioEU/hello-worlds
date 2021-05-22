package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/FlerioEU/hello-world/mongo-golang/db"
	"github.com/FlerioEU/hello-world/mongo-golang/routes"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var database db.Database
var r routes.Routes

func TestMain(m *testing.M) {
	// setup
	conf := db.DBConfig{
		Host:     "localhost",
		Port:     "27017",
		User:     "root",
		Password: "example",
		Database: "hw-golang",
	}

	database = db.New(conf)
	database.Connect()

	r = routes.New(database)
	// r.Run(":8080") => Blocks the thread :^).. sometimes you learn the hard way. Using executeRequest() instead

	// execute
	code := m.Run()

	// teardown
	database.DB.Collection("books").Drop(context.Background())
	database.Disconnect()

	os.Exit(code)
}

func TestPostBook(t *testing.T) {
	body := []byte(`{"title": "How to be testable", "author": "Mr. Test"}`)

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/books", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp := executeRequest(req)
	checkResponseCode(t, http.StatusCreated, resp.Code)

	var m map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &m)

	if m["Title"] != "How to be testable" {
		t.Errorf("Expected book title to be 'How to be testable'. Got '%v'", m["Title"])
	}

	if m["Author"] != "Mr. Test" {
		t.Errorf("Expected book title to be 'Mr. Test'. Got '%v'", m["Author"])
	}

	id := fmt.Sprintf("%v", m["Id"])
	if !primitive.IsValidObjectID(id) {
		t.Errorf("Expected valid object id. Got '%v'", id)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	r.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
