package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ToDOList struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"` //look this up
	Task   string             `json:"task,omitempty"`
	Status bool               `json:"status,omitempty"`
} //structure for storing todos
