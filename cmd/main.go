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

	cardRepository := database.NewCardRepository(cardCollection)
	cardService := service.NewCardService(cardRepository)
	cardController := controller.NewCardController(cardService)

	server := server.NewServer(engine, db, cardController)

	routes.SetupRouter(server)

	server.Run()
}
