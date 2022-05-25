package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/octaviomuller/deck-chips-server/database"
	"github.com/octaviomuller/deck-chips-server/helpers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SigninBody struct {
	UserId string `json:"userId"`
}

func SigninAnonymous(c *gin.Context) {
	var jsonData SigninBody

	err := c.BindJSON(&jsonData)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		c.Abort()
		return
	}

	userId, err := primitive.ObjectIDFromHex(jsonData.UserId)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		c.Abort()
		return
	}

	access, err := database.GetAnonymousAccess(userId)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		c.Abort()
		return
	}

	if access == nil {
		c.JSON(400, gin.H{"message": "User not found!"})
		c.Abort()
		return
	}

	user, err := database.GetUserById(userId)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		c.Abort()
		return
	}

	if user == nil {
		c.JSON(400, gin.H{"message": "User not found!"})
		c.Abort()
		return
	}

	token, err := helpers.CreateAccessToken(*user)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		c.Abort()
		return
	}
	fmt.Println(token)

	c.JSON(200, gin.H{"token": token})
}
