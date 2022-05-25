package database

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/octaviomuller/deck-chips-server/helpers"
	"github.com/octaviomuller/deck-chips-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

	var cardsList []models.Card

	findOptions, paginationErr := helpers.Pagination(c)
	if paginationErr != nil {
		return nil, paginationErr
	}

	fmt.Println(*findOptions)

	cur, findErr := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if findErr != nil {
		return nil, findErr
	}

	for cur.Next(context.TODO()) {
		var card models.Card

		decodeErr := cur.Decode(&card)
		if decodeErr != nil {
			return nil, decodeErr
		}

		cardsList = append(cardsList, card)
	}

	if findErr := cur.Err(); findErr != nil {
		return nil, findErr
	}

	cur.Close(context.TODO())

	return cardsList, nil
}
