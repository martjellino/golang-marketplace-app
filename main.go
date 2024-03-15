package main

import (
	"golang-marketplace-app/database"
	"golang-marketplace-app/router"
	"database/sql"
)

var (
	PORT = ":8000"
    DB  *sql.DB
)

func main() {
	database.StartDB()
	r := router.StartApp(DB)
	r.Run(PORT)
}
