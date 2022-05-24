package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/octaviomuller/deck-chips-server/controllers"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/cards/:code", controllers.GetCardByCode)
	r.GET("/cards", controllers.GetCards)

	r.Run()
}