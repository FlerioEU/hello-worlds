package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

type Database struct {
	DB     *mongo.Database
	ctx    context.Context
	cancel context.CancelFunc
	conf   DBConfig
	uri    string
}

func New(conf DBConfig) Database {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/", conf.User, conf.Password, conf.Host, conf.Port)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)

	return Database{
		uri:    uri,
		ctx:    ctx,
		cancel: cancel,
		conf:   conf,
	}
}

func (db *Database) Connect() {
	log.Printf("Connecting to MongoDB at '%v'\n", db.uri)
	client, err := mongo.Connect(db.ctx, options.Client().ApplyURI(db.uri))
	if err != nil {
		log.Fatalf("An error occured while connecting to the mongodb: %v", err)
	}

	db.DB = client.Database(db.conf.Database)
	log.Printf("Connected to MongoDB - Database '%v'", db.DB.Name())
}

func (db *Database) Disconnect() {
	log.Println("Disconnecting Database...")
	db.cancel()
	log.Println("Database disconnected")
}
