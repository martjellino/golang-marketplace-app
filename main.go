package main

import (
	"database/sql"
	"golang-marketplace-app/database"
	"golang-marketplace-app/router"

	"github.com/gin-gonic/gin"
)

var (
	PORT = ":8000"
	DB   *sql.DB
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	database.StartDB()
	r := router.StartApp()
	r.Run(PORT)
}
