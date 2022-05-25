package database

import (
	"context"

	"github.com/octaviomuller/deck-chips-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(user_info models.User) (interface{}, error) {
	collection := GetCollection("users")

	result, err := collection.InsertOne(context.TODO(), user_info)
	if err != nil {
		return nil, err
	}

	userId := result.InsertedID

	return userId, nil
}

func GetUserById(userId primitive.ObjectID) (*models.User, error) {
	collection := GetCollection("users")
	user := models.User{}

	filter := bson.D{{Key: "_id", Value: userId}}

	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
