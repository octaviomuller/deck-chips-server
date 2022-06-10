package database

import (
	"context"

	"github.com/octaviomuller/deck-chips-server/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type cardRepository struct {
	collection mongo.Collection
}

func NewCardRepository(collection mongo.Collection) *cardRepository {
	return &cardRepository{
		collection: collection,
	}
}

func (repository *cardRepository) FindOne(query interface{}, opts *options.FindOneOptions) (*models.Card, error) {
	result := &models.Card{}

	err := repository.collection.FindOne(context.TODO(), query, opts).Decode(result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (repository *cardRepository) FindMany(query interface{}, opts *options.FindOptions) (*[]models.Card, error) {
	result := &[]models.Card{}

	cur, findErr := repository.collection.Find(context.TODO(), query, opts)
	if findErr != nil {
		return nil, findErr
	}

	decodeErr := cur.All(context.TODO(), result)
	if decodeErr != nil {
		return nil, decodeErr
	}

	return result, nil
}
