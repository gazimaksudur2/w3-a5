package config

import (
	"os"

	"github.com/joho/godotenv"
)

func Load() {

	godotenv.Load()

}

func GoogleAPIKey() string {

	return os.Getenv("GOOGLE_API_KEY")

}

func TicketmasterAPIKey() string {

	return os.Getenv("TICKETMASTER_API_KEY")

}
