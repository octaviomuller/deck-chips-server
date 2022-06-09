package config

import (
	"context"
	"os"

	"github.com/octaviomuller/deck-chips-server/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDatabase() *mongo.Database {
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		panic("You must set your 'MONGODB_URI' environmental variable.")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	dbName := os.Getenv("DATABASE_NAME")
	if dbName == "" {
		errors.EnvVarError("DATABASE_NAME")
	}

	return client.Database(dbName)
}
