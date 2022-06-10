package database

import (
	"context"

	"github.com/octaviomuller/deck-chips-server/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type deckRepository struct {
	collection *mongo.Collection
}

func NewDeckRepository(collection mongo.Collection) *deckRepository {
	return &deckRepository{
		collection: &collection,
	}
}
func (repository *deckRepository) Insert(insertObj models.Deck) (*models.Deck, error) {
	_, err := repository.collection.InsertOne(context.TODO(), insertObj)
	if err != nil {
		return nil, err
	}

	return &insertObj, nil
}
