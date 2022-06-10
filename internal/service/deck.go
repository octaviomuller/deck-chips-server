package service

import (
	"errors"

	"github.com/octaviomuller/deck-chips-server/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type deckRepository interface {
	Insert(insertObj models.Deck) (*models.Deck, error)
	FindOne(query interface{}, opts *options.FindOneOptions) (*models.Deck, error)
	FindMany(query interface{}, opts *options.FindOptions) (*[]models.Deck, error)
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

func (service *deckService) GetDeckById(id string) (*models.DeckResponse, error) {
	objectId, objectIdErr := primitive.ObjectIDFromHex(id)
	if objectIdErr != nil {
		return nil, errors.New("Invalid objectId")
	}

	query := bson.M{
		"_id": objectId,
	}

	deck, err := service.deckRepository.FindOne(query, nil)
	if err != nil {
		return nil, err
	}

	deckResponse := models.DeckResponse{
		Id:            deck.Id,
		Title:         deck.Title,
		CoverCardCode: deck.CoverCardCode,
		CoverUrl:      deck.CoverUrl,
		Cards:         [40]models.Card{},
	}

	for i, cardCode := range deck.Cards {
		card, getCardErr := service.cardService.GetCardByCardCode(cardCode)
		if getCardErr != nil {
			return nil, errors.New("Card not found!")
		}

		deckResponse.Cards[i] = *card
	}

	return &deckResponse, nil
}

func (service *deckService) GetDecks() (*[]models.Deck, error) {
	query := bson.M{}
	opts := &options.FindOptions{}

	decks, err := service.deckRepository.FindMany(query, opts)
	if err != nil {
		return nil, err
	}

	return decks, err
}
