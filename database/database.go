package database

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

type Helpers interface {
	EnvVarError(variableName string)
}

func ConnectDatabase(envVarError func(variableName string)) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable.")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	dbName := os.Getenv("DATABASE_NAME")
	if dbName == "" {
		envVarError("DATABASE_NAME")
	}

	DB = client.Database(dbName)
}

func GetCollection(collectionName string) *mongo.Collection {
	return DB.Collection(collectionName)
}
