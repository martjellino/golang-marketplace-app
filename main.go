package main

import (
	"golang-marketplace-app/database"
	"golang-marketplace-app/router"
	"golang-marketplace-app/database"
)

var (
	PORT = ":8000"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(PORT)
}
