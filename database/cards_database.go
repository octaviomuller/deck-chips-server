package database

// import (
// 	"context"
// 	"fmt"

// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/mongo"
// )

// var collection *mongo.Collection = GetCollection("cards")

// func GetCards() {
// 	var cards bson.M

// 	query := bson.D{}

// 	err := collection.FindOne(context.TODO(), query).Decode(&cards)
// 	if err == nil {
// 		panic(err)
// 	}

// 	fmt.Println(cards)
// }
