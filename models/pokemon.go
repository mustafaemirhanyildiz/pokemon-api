package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Pokemon struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name   string             `json:"name,omitempty" bson:"name,omitempty"`
	Type   string             `json:"type,omitempty" bson:"type,omitempty"`
	Number int                `json:"number,omitempty" bson:"number,omitempty"`
}
