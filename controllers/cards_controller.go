package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/octaviomuller/deck-chips-server/database"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetCard(c *gin.Context) {
	code := c.Params.ByName("code")

	card, err := database.GetCard(code)
	fmt.Println(card, err)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Card not found!", // TODO: Enumerate errors
			})
			return
		}

		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, card)
}
