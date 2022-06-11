package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/octaviomuller/deck-chips-server/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type deckService interface {
	CreateDeck(title string, coverCardCode string, cards [40]string) (*models.Deck, error)
	GetDeckById(id string) (*models.DeckResponse, error)
	GetDecks() (*[]models.Deck, error)
	UpdateDeck(id string, title *string, coverCardCode *string, cards *[]string) error
}

type DeckController struct {
	deckService deckService
}

func NewDeckController(service deckService) *DeckController {
	return &DeckController{
		deckService: service,
	}
}

func (controller *DeckController) Post(context *gin.Context) {
	body := models.CreateDeck{}

	bodyErr := context.BindJSON(&body)
	if bodyErr != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	title := body.Title
	coverCardCode := body.CoverCardCode
	cards := body.Cards

	deck, serviceErr := controller.deckService.CreateDeck(title, coverCardCode, cards)
	if serviceErr != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": serviceErr.Error(),
		})
	}

	context.JSON(http.StatusOK, deck)
	return
}

func (controller *DeckController) Get(context *gin.Context) {
	id := context.Params.ByName("id")
	if id == "" {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Invalid id",
		})
		return
	}

	deck, serviceErr := controller.deckService.GetDeckById(id)
	if serviceErr != nil {
		if serviceErr == mongo.ErrNoDocuments {
			context.JSON(http.StatusBadRequest, gin.H{
				"message": "Deck not found!",
			})
			return
		}

		context.JSON(http.StatusInternalServerError, serviceErr)
		return
	}

	context.JSON(http.StatusOK, deck)
	return
}

func (controller *DeckController) Index(context *gin.Context) {
	decks, serviceError := controller.deckService.GetDecks()
	if serviceError != nil {
		if serviceError == mongo.ErrNoDocuments {
			context.JSON(http.StatusBadRequest, gin.H{
				"message": "Decks not found!",
			})
			return
		}

		context.JSON(http.StatusInternalServerError, gin.H{
			"message": serviceError.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, decks)
	return
}

func (controller *DeckController) Patch(context *gin.Context) {
	id := context.Params.ByName("id")
	if id == "" {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Invalid id",
		})
		return
	}

	body := models.UpdateDeck{}

	bodyErr := context.BindJSON(&body)
	if bodyErr != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	title := body.Title
	coverCardCode := body.CoverCardCode
	cards := body.Cards

	serviceErr := controller.deckService.UpdateDeck(id, title, coverCardCode, cards)
	if serviceErr != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": serviceErr.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Update deck successfully",
	})
	return
}
