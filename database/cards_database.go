package database

import (
	"context"

	"github.com/octaviomuller/deck-chips-server/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetCard() models.Card {
	collection := GetCollection("cards")
	card := models.Card{}

	query := bson.D{}

	err := collection.FindOne(context.TODO(), query).Decode(&card)
	if err != nil {
		panic(err)
	}

	return card
}
