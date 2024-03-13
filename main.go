package main

import (
	"golang-marketplace-app/database"
	"golang-marketplace-app/router"
	"log"
)

var (
	PORT = ":8000"
)

func main() {
	db, err := database.InitDB("postgres://postgres:P4ssW0rd@localhost:5434/marketplace_db?sslmode=disable")
	if err != nil {
		log.Fatal("Error initializing database:", err)
	}

	r := router.StartApp(db)
	r.Run(PORT)
}
