package main

import (
	"todo-api/api"

	log "github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not coming through %v", err)
	}
	log.Info("Ok Boss")

	api.Run()
}
