package main

import (
	"github.com/MatthewJamesBoyle/bindates/parser"
	"github.com/joho/godotenv"
	"log"
	"os"
	"fmt"
)

func main() {
	errs := godotenv.Load()
	if errs != nil {
		log.Fatal("Error loading .env file")
	}

	dates := parser.Parse(parser.Config{
		os.Getenv("URL"),
		os.Getenv("HOUSE_NUMBER"),
		os.Getenv("POSTCODE"),
	})

	for key, value := range dates {
		fmt.Println(key, value)
	}
}
