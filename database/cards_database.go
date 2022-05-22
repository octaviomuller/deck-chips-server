package database

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/octaviomuller/deck-chips-server/helpers"
	"github.com/octaviomuller/deck-chips-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetCardByCode(code string) (models.Card, error) {
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

func GetCards(c *gin.Context) ([]models.Card, error) {
	collection := GetCollection("cards")
	cards := []models.Card{}
	findOptions := options.FindOptions{}

	helpers.Pagination(c, &findOptions)

	cur, err := collection.Find(context.TODO(), bson.M{}, &findOptions)

	if err != nil {
		return cards, err
	}

	cur.Decode(&cards)
	return cards, nil
}
