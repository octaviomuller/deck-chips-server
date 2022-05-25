package main

import (
	"github.com/octaviomuller/deck-chips-server/database"
	"github.com/octaviomuller/deck-chips-server/helpers"
	"github.com/octaviomuller/deck-chips-server/routes"
)

func main() {
	database.ConnectDatabase(helpers.EnvVarError)

	routes.HandleRequests()
}
