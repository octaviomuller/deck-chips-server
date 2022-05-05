package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/octaviomuller/deck-chips-server/helpers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDatabase() {
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
		helpers.EnvVarError("DATABASE_NAME")
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
}

func GetCards() {
	var cards bson.M

	query := bson.D{}

	err := DB.Collection("cards").FindOne(context.TODO(), query).Decode(&cards)
	if err == nil {
		panic(err)
	}

	fmt.Println(cards)
}
