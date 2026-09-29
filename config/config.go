package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Load() {

	err := godotenv.Load()

	if err != nil {
		log.Println("No .env file found, using system environment")
	}

}

func GoogleAPIKey() string {

	return os.Getenv("GOOGLE_API_KEY")

}

func TicketmasterAPIKey() string {

	return os.Getenv("TICKETMASTER_API_KEY")

}
