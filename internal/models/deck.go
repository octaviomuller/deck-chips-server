package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Deck struct {
	Id            primitive.ObjectID `bson:"_id" json:"_id"`
	Title         string             `json:"title"`
	CoverCardCode string             `json:"coverCardCode" bson:"coverCardCode"`
	CoverUrl      string             `json:"coverUrl" bson:"coverUrl"`
	Cards         [40]string         `json:"cards"`
}

type DeckResponse struct {
	Id            primitive.ObjectID `json:"_id"`
	Title         string             `json:"title"`
	CoverCardCode string             `json:"coverCardCode"`
	CoverUrl      string             `json:"coverUrl"`
	Cards         [40]Card           `json:"cards"`
}

type CreateDeck struct {
	Title         string     `json:"title"`
	CoverCardCode string     `json:"coverCardCode"`
	Cards         [40]string `jsons:"cards"`
}

type UpdateDeck struct {
	Title         *string   `json:"title"`
	CoverCardCode *string   `json:"coverCardCode" bson:"coverCardCode"`
	CoverUrl      *string   `json:"coverUrl" bson:"coverUrl"`
	Cards         *[]string `jsons:"cards"`
}
