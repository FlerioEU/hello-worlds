package models

type Book struct {
	Id     string `bson:"_id,omitempty"`
	Title  string `bson:"title,omitempty"`
	Author string `bson:"author,omitempty"`
}
