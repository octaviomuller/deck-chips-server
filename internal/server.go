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
}

func NewServer(engine *gin.Engine, database *mongo.Database, cardController *controller.CardController) *Server {
	return &Server{
		Engine:         engine,
		Database:       database,
		CardController: cardController,
	}
}

func (server *Server) Run() {
	server.Engine.Run()
}
