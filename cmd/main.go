package main

import (
	"github.com/gin-gonic/gin"
	"github.com/octaviomuller/deck-chips-server/config"
	server "github.com/octaviomuller/deck-chips-server/internal"
	"github.com/octaviomuller/deck-chips-server/internal/controller"
	"github.com/octaviomuller/deck-chips-server/internal/database"
	"github.com/octaviomuller/deck-chips-server/internal/routes"
	"github.com/octaviomuller/deck-chips-server/internal/service"
)

func main() {
	engine := gin.Default()

	db := config.ConnectDatabase()

	cardCollection := db.Collection("cards")
	deckCollection := db.Collection("decks")

	cardRepository := database.NewCardRepository(*cardCollection)
	deckRepository := database.NewDeckRepository(*deckCollection)

	cardService := service.NewCardService(cardRepository)
	deckService := service.NewDeckService(deckRepository, *cardService)

	cardController := controller.NewCardController(cardService)
	deckController := controller.NewDeckController(deckService)

	server := server.NewServer(engine, db, cardController, deckController)

	routes.SetupRouter(server)

	server.Run()
}
