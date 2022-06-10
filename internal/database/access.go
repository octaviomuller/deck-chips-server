package database

import (
	"context"

	"github.com/octaviomuller/deck-chips-server/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type accessRepository struct {
	collection *mongo.Collection
}

func NewAccessRepository(collection *mongo.Collection) *accessRepository {
	return &accessRepository{
		collection: collection,
	}
}

func (repository accessRepository) InsertOne(access_info models.Access) (*mongo.InsertOneResult, error) {
	result, err := repository.collection.InsertOne(context.TODO(), access_info)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (repository accessRepository) FindOne(query interface{}) (*models.Access, error) {
	result := &models.Access{}

	err := repository.collection.FindOne(context.TODO(), query).Decode(result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
