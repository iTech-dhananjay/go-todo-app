package models

import "go.mongodb.org/mongo-driver/mongo"

type DB struct {
	Client *mongo.Client
}

type Todo struct {
	ID     string `json:"id" bson:"_id,omitempty"`
	Title  string `json:"title" bson:"title"`
	Status string `json:"status" bson:"status"`
}
