package database

import (
	"context"

	"github.com/octaviomuller/deck-chips-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateAccess(access_info models.Access) interface{} {
	collection := GetCollection("accesses")

	result, err := collection.InsertOne(context.TODO(), access_info)
	if err != nil {
		panic(err)
	}

	userId := result.InsertedID

	return userId
}

func GetAnonymousAccess(userId primitive.ObjectID) (*models.Access, error) {
	collection := GetCollection("accesses")
	access := models.Access{}

	query := bson.D{
		{Key: "userId", Value: userId},
		{Key: "active", Value: true},
		{Key: "method", Value: "anonymous"},
	}

	err := collection.FindOne(context.TODO(), query).Decode(&access)
	if err != nil {
		return nil, err
	}

	return &access, nil
}
