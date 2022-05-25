package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/octaviomuller/deck-chips-server/database"
	"github.com/octaviomuller/deck-chips-server/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SignupWithEmailAndPassword(c *gin.Context) {
}

func SignupWithGoogle(c *gin.Context) {
}

func SignupWithFacebook(c *gin.Context) {
}

func SignupAnonymous(c *gin.Context) {
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
