package main

import (
	"github.com/octaviomuller/deck-chips-server/database"
)

func main() {
	database.ConnectDatabase()
	database.GetCards()
}
