package main

import (
	"fmt"
	"log"

	"rest-api-boilerplate/common"

	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	fmt.Println("te")
	app := App{}
	app.Initialize()
	defer common.CloseDB()
	app.Run()
}
