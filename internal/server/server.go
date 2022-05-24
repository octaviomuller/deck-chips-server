package server

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	Database *mongo.Database
	Router   *gin.RouterGroup
}

func NewServer(database *mongo.Database, router gin.RouterGroup) *Server {

}

func (server *Server) HandleRequests() {

}
