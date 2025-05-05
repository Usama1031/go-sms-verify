package api

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func envAccountSID() string {
	println(godotenv.Unmarshal(".env"))
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln(err)
		log.Fatal("Error Loading .env file!")
	}
	return os.Getenv("TWILIO_ACCOUNT_SID")

}

func envAuthToken() string {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error Loading .env file!")
	}
	return os.Getenv("TWILIO_AUTH_TOKEN")
}

func envServicesSID() string {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error Loading .env file!")
	}
	return os.Getenv("TWILIO_SERVICES_SID")
}
