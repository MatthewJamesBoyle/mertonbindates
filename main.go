package main

import (
	"os"
	"github.com/joho/godotenv"
	"log"
	"github.com/MatthewJamesBoyle/bindates/parser"
)

func main() {
	errs := godotenv.Load()
	if errs != nil {
		log.Fatal("Error loading .env file")
	}

	parser.Parse(parser.Config{
		os.Getenv("URL"),
		os.Getenv("HOUSE_NUMBER"),
		os.Getenv("POSTCODE"),
	})
}
