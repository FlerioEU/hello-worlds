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
	"github.com/FlerioEU/hello-world/mongo-golang/models"
	"github.com/FlerioEU/hello-world/mongo-golang/routes"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var id string
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

	// TODO
	// would be better to preload some objects into the collection with some preset ids to have an independant testing flow
	// now everything is dependant on the POST working correctly and in the right order

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

	id = fmt.Sprintf("%v", m["Id"])
	if !primitive.IsValidObjectID(id) {
		t.Errorf("Expected valid object id. Got '%v'", id)
	}
}

func TestGetBooks(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/books", nil)

	resp := executeRequest(req)
	checkResponseCode(t, http.StatusOK, resp.Code)

	var b []models.Book
	json.Unmarshal(resp.Body.Bytes(), &b)

	if len(b) != 1 {
		t.Errorf("Expected one book in slice but got %v!", len(b))
	}

	if b[0].Author != "Mr. Test" {
		t.Errorf("Expected book title to be 'Mr. Test'. Got '%v'", b[0].Author)
	}

	if b[0].Title != "How to be testable" {
		t.Errorf("Expected book title to be 'How to be testable'. Got '%v'", b[0].Title)
	}

	if !primitive.IsValidObjectID(b[0].Id) {
		t.Errorf("Expected valid object id. Got '%v'", b[0].Id)
	}
}

func TestGetBook(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/books/"+id, nil)

	resp := executeRequest(req)
	checkResponseCode(t, http.StatusOK, resp.Code)

	var b models.Book
	json.Unmarshal(resp.Body.Bytes(), &b)

	if b.Author != "Mr. Test" {
		t.Errorf("Expected book title to be 'Mr. Test'. Got '%v'", b.Author)
	}

	if b.Title != "How to be testable" {
		t.Errorf("Expected book title to be 'How to be testable'. Got '%v'", b.Title)
	}

	if !primitive.IsValidObjectID(b.Id) {
		t.Errorf("Expected valid object id. Got '%v'", b.Id)
	}
}

func TestUpdateBook(t *testing.T) {
	body := []byte(`{"title": "Updated the shit out of you :)", "author": "Mrs. Updater"}`)

	req, _ := http.NewRequest(http.MethodPut, "/api/v1/books/"+id, bytes.NewBuffer(body))

	resp := executeRequest(req)
	checkResponseCode(t, http.StatusOK, resp.Code)

	var b models.Book
	json.Unmarshal(resp.Body.Bytes(), &b)

	if b.Author != "Mrs. Updater" {
		t.Errorf("Expected book title to be 'Mrs. Updater'. Got '%v'", b.Author)
	}

	if b.Title != "Updated the shit out of you :)" {
		t.Errorf("Expected book title to be 'Updated the shit out of you :)'. Got '%v'", b.Title)
	}

	if !primitive.IsValidObjectID(b.Id) {
		t.Errorf("Expected valid object id. Got '%v'", b.Id)
	}

	if b.Id != id {
		t.Errorf("Expected unchanged id but has been changed. Got '%v' but expected '%v'", b.Id, id)
	}
}

func TestDeleteBook(t *testing.T) {
	reqDel, _ := http.NewRequest(http.MethodDelete, "/api/v1/books/"+id, nil)

	respDel := executeRequest(reqDel)
	checkResponseCode(t, http.StatusOK, respDel.Code)

	reqGet, _ := http.NewRequest(http.MethodGet, "/api/v1/books/"+id, nil)

	respGet := executeRequest(reqGet)
	checkResponseCode(t, http.StatusNotFound, respGet.Code)
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
