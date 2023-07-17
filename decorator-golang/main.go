package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type Database struct {
}

type Person struct {
	Name       string
	Age        int
	Occupation string
}

func (d Database) Get(id string) Person {
	return Person{
		Name:       "Leeroy",
		Age:        43,
		Occupation: "Plumber",
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/person/", personHandler("/api/v1/person/", Database{}))
	mux.HandleFunc("/api/v1/person", personsHandler(Database{}))

	log.Printf("listening on port 3000...")

	http.ListenAndServe(":3000", mux)
}

// curl localhost:3000/api/v1/person/{id}
func personHandler(path string, db Database) http.HandlerFunc {
	// needs the path to get the path param e.g. /hello-world/{id}
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("request: %s - %s", r.Method, r.URL.Path)

		id := strings.TrimPrefix(r.URL.Path, path)
		log.Printf("incoming request with id '%s'", id)

		p := db.Get(id)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(p)
	}
}

// curl localhost:3000/api/v1/person
func personsHandler(db Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("request: %s - %s", r.Method, r.URL.Path)

		var ps []Person
		p := db.Get("")
		ps = append(ps, p)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ps)
	}
}
