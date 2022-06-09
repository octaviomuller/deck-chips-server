package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/octaviomuller/deck-chips-server/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type cardService interface {
	GetCardByCardCode(cardCode string) (*models.Card, error)
	GetCards(
		page string,
		limit string,
		name string,
		region string,
		cost string,
		cardType string,
		rarity string,
		set string,
	) (*[]models.Card, error)
}

type CardController struct {
	cardService cardService
}

func NewCardController(service cardService) *CardController {
	return &CardController{
		cardService: service,
	}
}

func (controller *CardController) Get(context *gin.Context) {
	cardCode := context.Params.ByName("code")
	if cardCode == "" { // TODO: validation package
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Invalid card code",
		})
		return
	}

	card, serviceErr := controller.cardService.GetCardByCardCode(cardCode)
	if serviceErr != nil {
		if serviceErr == mongo.ErrNoDocuments {
			context.JSON(http.StatusBadRequest, gin.H{
				"message": "Card not found!", // TODO: Enumerate errors + error package
			})
			return
		}

		context.JSON(http.StatusInternalServerError, serviceErr) // TODO: Response package
		return
	}

	context.JSON(http.StatusOK, card)
	return
}

func (controller *CardController) Index(context *gin.Context) {
	page := context.Request.URL.Query().Get("page")
	limit := context.Request.URL.Query().Get("limit")

	name := context.Request.URL.Query().Get("name")
	region := context.Request.URL.Query().Get("region")
	cost := context.Request.URL.Query().Get("cost")
	cardType := context.Request.URL.Query().Get("type")
	rarity := context.Request.URL.Query().Get("rarity")
	set := context.Request.URL.Query().Get("set")

	cards, err := controller.cardService.GetCards(
		page,
		limit,
		name,
		region,
		cost,
		cardType,
		rarity,
		set,
	)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			context.JSON(http.StatusBadRequest, gin.H{
				"message": "Cards not found!",
			})
			return
		}

		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, cards)
	return
}
