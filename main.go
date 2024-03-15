package main

import (
	"database/sql"
	"golang-marketplace-app/database"
	"golang-marketplace-app/router"
)

var (
	PORT = ":8000"
	DB   *sql.DB
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(PORT)
}
