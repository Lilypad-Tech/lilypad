package main

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"

	"github.com/bacalhau-project/lilypad/cmd/lilypad"
)

func main() {
	lilypad.Execute()
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Debug().Msgf(".env not found")
	}

}
