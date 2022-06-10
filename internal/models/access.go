package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Access struct {
	Method      string             `json:"method" bson:"method"`
	Value       string             `json:"value" bson:"value,omitempty"`
	UsePassword bool               `json:"use_password" bson:"usePassword"`
	Active      bool               `json:"active" bson:"active"`
	UserId      primitive.ObjectID `json:"userId" bson:"userId,omitempty"`
}
