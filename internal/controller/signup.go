package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/octaviomuller/deck-chips-server/database"
	"github.com/octaviomuller/deck-chips-server/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func EmailAndPassword(c *gin.Context) {
}

func Google(c *gin.Context) {
}

func Facebook(c *gin.Context) {
}

func Anonymous(c *gin.Context) {
	userId := primitive.NewObjectID()
	_, err := database.CreateUser(models.User{
		Id: userId,
	})
	if err != nil {
		c.JSON(500, err)
	}

	database.CreateAccess(models.Access{
		Method:      "anonymous",
		Value:       userId.Hex(),
		UsePassword: false,
		Active:      true,
		UserId:      userId,
	})

	c.JSON(200, userId)
}
