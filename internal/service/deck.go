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
	UpdateById(id primitive.ObjectID, updateObj interface{}) error
	DeleteOne(query interface{}) error
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

func (service *deckService) CreateDeck(title string, coverCardCode string, cards []string, userName string, userId string) (*models.Deck, error) {
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
		UserName:      userName,
		UserId:        userId,
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
		Cards:         []models.Card{},
		UserName:      deck.UserName,
		UserId:        deck.UserId,
	}

	for _, cardCode := range deck.Cards {
		card, getCardErr := service.cardService.GetCardByCardCode(cardCode)
		if getCardErr != nil {
			return nil, errors.New("Card not found!")
		}

		deckResponse.Cards = append(deckResponse.Cards, *card)
	}

	return &deckResponse, nil
}

func (service *deckService) GetDecks(title string, userId string) (*[]models.Deck, error) {
	query := bson.M{}
	opts := &options.FindOptions{}

	if title != "" {
		query["title"] = bson.M{
			"$regex":   title,
			"$options": "i",
		}
	}
	if userId != "" {
		query["userId"] = userId
	}

	decks, err := service.deckRepository.FindMany(query, opts)
	if err != nil {
		return nil, err
	}

	return decks, err
}

func (service *deckService) UpdateDeck(id string, title *string, coverCardCode *string, cards *[]string) error {
	objectId, objectIdErr := primitive.ObjectIDFromHex(id)
	if objectIdErr != nil {
		return errors.New("Invalid objectId")
	}

	update := bson.M{}

	if title != nil {
		update["title"] = title
	}

	if coverCardCode != nil {
		card, gerCardErr := service.cardService.GetCardByCardCode(*coverCardCode)
		if gerCardErr != nil {
			return errors.New("Invalid card code")
		}

		update["coverCardCode"] = coverCardCode
		update["coverUrl"] = card.Assets[0].FullAbsolutePath
	}

	if cards != nil {
		update["cards"] = cards
	}

	updateObj := bson.M{"$set": update}

	updateErr := service.deckRepository.UpdateById(objectId, updateObj)
	if updateErr != nil {
		return updateErr
	}

	return nil
}

func (service *deckService) DeleteDeck(id string) error {
	objectId, objectIdErr := primitive.ObjectIDFromHex(id)
	if objectIdErr != nil {
		return errors.New("Invalid objectId")
	}

	query := bson.M{
		"_id": objectId,
	}

	err := service.deckRepository.DeleteOne(query)
	if err != nil {
		return err
	}

	return nil
}
