package routes

import (
	server "github.com/octaviomuller/deck-chips-server/internal"
)

func SetupRouter(server *server.Server) {
	router := server.Engine.Group("/api/v1")
	{
		signup := router.Group("/signup")
		{
			signup.POST("/signup/anonymous", server.SignupController.Anonymous)
		}

		signin := router.Group("/signin")
		{
			signin.POST("/signin/anonymous", server.SigninController.Anonymous)
		}

		router.Use(middleware.Authenticate)
		auth := router.Group("/auth")
		{
			card := auth.Group("/cards")
			{
				card.GET("/:code", server.CardController.Get)
				card.GET("/", server.CardController.Index)
			}
		}

	}
}
