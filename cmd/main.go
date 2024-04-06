package main

import (
	"log"

	"github.com/ilhamgepe/todos-backend/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	s := server.NewServer()

	if err := s.ListenAndServe(); err != nil {
		log.Fatal("Error starting server: ", err)
	}

}
