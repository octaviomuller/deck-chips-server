package database

import (
	"context"

	"github.com/octaviomuller/deck-chips-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetCard(code string) (models.Card, error) {
	collection := GetCollection("cards")
	card := models.Card{}

	query := bson.M{
		"cardCode": code,
	}

	err := collection.FindOne(context.TODO(), query).Decode(&card)
	if err == mongo.ErrNoDocuments {
		return card, mongo.ErrNoDocuments
	}

	return card, nil
}
