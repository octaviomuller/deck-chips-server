package server

import (
	"github.com/gin-gonic/gin"
	"github.com/octaviomuller/deck-chips-server/internal/controller"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	Engine         *gin.Engine
	Database       *mongo.Database
	CardController *controller.CardController
	DeckController *controller.DeckController
}

func NewServer(engine *gin.Engine, database *mongo.Database, cardController *controller.CardController, deckController *controller.DeckController) *Server {
	return &Server{
		Engine:         engine,
		Database:       database,
		CardController: cardController,
		DeckController: deckController,
	}
}

func (server *Server) Run() {
	server.Engine.Run()
}
