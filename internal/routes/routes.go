package routes

import (
	server "github.com/octaviomuller/deck-chips-server/internal"
)

func SetupRouter(server *server.Server) {
	router := server.Engine.Group("/api/v1")
	{
		card := router.Group("/cards")
		{
			card.GET("/:code", server.CardController.Get)
			card.GET("/", server.CardController.Index)
		}
		deck := router.Group("/decks")
		{
			deck.POST("/", server.DeckController.Post)
			deck.GET("/:id", server.DeckController.Get)
			deck.GET("/", server.DeckController.Index)
			deck.PATCH("/:id", server.DeckController.Patch)
			deck.DELETE("/:id", server.DeckController.Delete)
		}
	}
}
