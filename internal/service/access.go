package service

import (
	"github.com/octaviomuller/deck-chips-server/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type accessRepository interface {
	InsertOne(access_info models.Access) (*mongo.InsertOneResult, error)
	FindOne(query interface{}) (*models.Access, error)
}

type accessService struct {
	accessRepository accessRepository
}

func NewAccessService(repository accessRepository) *accessService {
	return &accessService{
		accessRepository: repository,
	}
}

func (service *accessService) CreateAnonymous() (interface{}, error) {
	result, err := service.accessRepository.InsertOne(models.Access{Method: "anonymous"})
	if err != nil {
		return nil, err
	}

	return result.InsertedID, nil
}

func (service *accessService) FindAnonymousAccess(userId interface{}) (*models.Access, error) {
	query := bson.D{
		{Key: "userId", Value: userId},
		{Key: "active", Value: true},
		{Key: "method", Value: "anonymous"},
	}

	access, err := service.accessRepository.FindOne(query)
	if err != nil {
		return nil, err
	}

	return access, nil
}
