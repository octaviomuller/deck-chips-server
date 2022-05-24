package helpers

import (
	"log"

	"github.com/gin-gonic/gin"
)

func EnvVarError(variableName string) {
	log.Fatal("You must set your '" + variableName + "' environmental variable.")
}

func ResponseMessage(message string) gin.H {
	return gin.H{
		"message": message,
	}
}
