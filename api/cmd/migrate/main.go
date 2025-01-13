package main

import (
	"github.com/azeezlala/assessment/api/cmd/cli"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	if os.Getenv("ENV") != "PRODUCTION" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("loading env error: %v", err)
		}
	}

	cli.Execute()

}
