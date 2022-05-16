package main

import (
	"github.com/joho/godotenv"
	server2 "github.com/viralgame/server/handlers"
	"log"
)

func main() {
	// Load the .env file in the current directory
	err := godotenv.Load()
	if err != nil {
		log.Fatal("unable to load env : ", err)
	}
	server2.Handler.Run()
}