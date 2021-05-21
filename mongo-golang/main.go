package main

import (
	"log"
	"os"

	"github.com/FlerioEU/hello-world/mongo-golang/db"
	"github.com/FlerioEU/hello-world/mongo-golang/routes"
)

func main() {
	log.Println("Starting service...")

	conf := db.DBConfig{
		User:     os.Getenv("MONGO_USER"),
		Password: os.Getenv("MONGO_PASSWORD"),
		Host:     os.Getenv("MONGO_HOST"),
		Port:     os.Getenv("MONGO_PORT"),
		Database: os.Getenv("MONGO_DATABASE"),
	}

	db := db.New(conf)
	db.Connect()
	defer db.Disconnect()

	r := routes.New(db)

	r.Run(":8080")
}
