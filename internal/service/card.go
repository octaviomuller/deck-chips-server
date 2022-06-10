package service

import (
	"errors"
	"strconv"
	"strings"

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

func (service *cardService) GetCards(
	page string,
	limit string,
	name string,
	region string,
	cost string,
	cardType string,
	rarity string,
	set string,
) (*[]models.Card, error) {
	query := bson.M{
		"collectible": true,
	}

	if name != "" {
		query["name"] = bson.M{
			"$regex":   name,
			"$options": "i",
		}
	}
	if region != "" {
		query["regions"] = bson.M{
			"$regex":   region,
			"$options": "i",
		}
	}
	if cost != "" {
		costNumber, err := strconv.Atoi(cost)
		if err != nil {
			return nil, errors.New("Invalid type for variable 'cost'")
		}

		query["cost"] = costNumber
	}
	if cardType != "" {
		if cardTypeLower := strings.ToLower(cardType); cardTypeLower == "champion" {
			query["supertype"] = bson.M{
				"$regex":   cardType,
				"$options": "i",
			}
		} else {
			query["type"] = bson.M{
				"$regex":   cardType,
				"$options": "i",
			}
		}
	}
	if rarity != "" {
		query["rarity"] = bson.M{
			"$regex":   rarity,
			"$options": "i",
		}
	}
	if set != "" {
		query["set"] = bson.M{
			"$regex":   set,
			"$options": "i",
		}
	}

	opts, paginationErr := helper.Pagination(page, limit)
	if paginationErr != nil {
		return nil, paginationErr
	}

	opts.SetSort(bson.D{
		{Key: "cost", Value: 1},
		{Key: "_id", Value: 1},
	})

	cards, err := service.cardRepository.FindMany(query, opts)
	if err != nil {
		return nil, err
	}

	return cards, nil
}
