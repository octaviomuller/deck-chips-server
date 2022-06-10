package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id        primitive.ObjectID `json:"_id" bson:"_id"`
	Username  string             `json:"username" bson:"username,omitempty"`
	Password  string             `json:"password" bson:"password,omitempty"`
	Email     string             `json:"email" bson:"email,omitempty"`
	Cellphone string             `json:"cellphone" bson:"cellphone,omitempty"`
}
