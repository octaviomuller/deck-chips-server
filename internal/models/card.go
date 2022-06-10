package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type AssetsPath struct {
	GameAbsolutePath string `json:"gameAbsolutePath"`
	FullAbsolutePath string `json:"fullAbsolutePath"`
}

type Card struct {
	Id                    primitive.ObjectID `bson:"_id" json:"_id"`
	AssociatedCards       []string           `json:"associatedCards"`
	AssociatedCardRefs    []string           `json:"associatedCardsRefs"`
	Assets                []AssetsPath       `json:"assets"`
	Regions               []string           `json:"regions"`
	RegionRefs            []string           `json:"regionRef"`
	Attack                int                `json:"attack"`
	Cost                  int                `json:"cost"`
	Health                int                `json:"health"`
	Description           string             `json:"description"`
	DescriptionRaw        string             `json:"descriptionRaw"`
	LevelupDescription    string             `json:"levelupDescription"`
	LevelupDescriptionRaw string             `json:"levelupDescriptionRaw"`
	FlavorText            string             `json:"flavorText"`
	ArtistName            string             `json:"artistName"`
	Name                  string             `json:"name"`
	CardCode              string             `json:"cardCode"`
	Keywords              []string           `json:"keywords"`
	KeywordRefs           []string           `json:"keywordRefs"`
	SpellSpeed            string             `json:"spellSpeed"`
	SpellSpeedRef         string             `json:"spellSpeedRef"`
	Rarity                string             `json:"rarity"`
	RarityRef             string             `json:"rarityRef"`
	Subtypes              []string           `json:"subtypes"`
	Supertype             string             `json:"supertype"`
	Type                  string             `json:"type"`
	Collectible           bool               `json:"collectible"`
	Set                   string             `json:"set"`
}
