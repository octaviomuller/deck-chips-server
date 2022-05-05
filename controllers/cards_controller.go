package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/octaviomuller/deck-chips-server/database"
)

func GetCard(c *gin.Context) {

	card := database.GetCard()

	c.JSON(200, card)
}
