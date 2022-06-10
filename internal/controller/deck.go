package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/octaviomuller/deck-chips-server/internal/models"
)

type deckService interface {
	CreateDeck(title string, coverCardCode string, cards [40]string) (*models.Deck, error)
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
