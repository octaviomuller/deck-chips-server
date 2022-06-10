package service

import (
	"github.com/octaviomuller/deck-chips-server/internal/models"
	"github.com/octaviomuller/deck-chips-server/pkg/helper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type cardRepository interface {
	FindOne(query interface{}, opts *options.FindOneOptions) (*models.Card, error)
	FindMany(query interface{}, opts *options.FindOptions) (*[]models.Card, error)
}

type cardService struct {
	cardRepository cardRepository
}

func NewCardService(repository cardRepository) *cardService {
	return &cardService{
		cardRepository: repository,
	}
}

func (service *cardService) GetCardByCardCode(cardCode string) (*models.Card, error) {
	query := bson.M{
		"cardCode": cardCode,
	}

	card, err := service.cardRepository.FindOne(query, nil)
	if err != nil {
		return nil, err
	}

	return card, nil
}

func (service *cardService) GetCards(page string, limit string) (*[]models.Card, error) {
	query := bson.M{}
	opts, paginationErr := helper.Pagination(page, limit)
	if paginationErr != nil {
		return nil, paginationErr
	}

	cards, err := service.cardRepository.FindMany(query, opts)
	if err != nil {
		return nil, err
	}

	return cards, nil
}
