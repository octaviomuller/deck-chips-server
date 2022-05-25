package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/octaviomuller/deck-chips-server/controllers"
	"github.com/octaviomuller/deck-chips-server/middlewares"
)

func HandleRequests() {
	r := gin.Default()

	r.POST("/signup/anonymous", controllers.SignupAnonymous)
	r.POST("/signin/anonymous", controllers.SigninAnonymous)

	r.Use(middlewares.Authenticate)
	r.GET("/cards", controllers.GetCard)

	r.Run()
}
