package service

import (
	"errors"

	"github.com/octaviomuller/deck-chips-server/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type deckRepository interface {
	Insert(insertObj models.Deck) (*models.Deck, error)
}

type deckService struct {
	deckRepository deckRepository
	cardService    cardService
}

func NewDeckService(deckRepository deckRepository, cardService cardService) *deckService {
	return &deckService{
		deckRepository: deckRepository,
		cardService:    cardService,
	}
}

func (service *deckService) CreateDeck(title string, coverCardCode string, cards [40]string) (*models.Deck, error) {
	card, gerCardErr := service.cardService.GetCardByCardCode(coverCardCode)
	if gerCardErr != nil {
		return nil, errors.New("Invalid card code")
	}

	insertDeck := models.Deck{
		Id:            primitive.NewObjectID(),
		Title:         title,
		CoverCardCode: coverCardCode,
		CoverUrl:      card.Assets[0].FullAbsolutePath,
		Cards:         cards,
	}

	deck, insertErr := service.deckRepository.Insert(insertDeck)
	if insertErr != nil {
		return nil, insertErr
	}

	return deck, nil
}
