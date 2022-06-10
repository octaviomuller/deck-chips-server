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
		}
	}
}
