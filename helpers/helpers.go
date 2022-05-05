package helpers

import "log"

func EnvVarError(variableName string) {
	log.Fatal("You must set your '" + variableName + "' environmental variable.")
}
