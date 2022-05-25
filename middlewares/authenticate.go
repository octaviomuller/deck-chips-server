package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/octaviomuller/deck-chips-server/helpers"
)

func Authenticate(c *gin.Context) {
	const BEARER_SCHEMA = "Bearer "
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.JSON(401, gin.H{"message": "Token is missing!"})
		c.Abort()
		return
	}

	tokenString := authHeader[len(BEARER_SCHEMA):]

	user, err := helpers.ValidateAccessToken(tokenString)
	if err != nil {
		fmt.Println(err)
		c.JSON(401, gin.H{"message": "Invalid token!"})
		c.Abort()
		return
	}

	c.Set("user", user)

	c.Next()
}
