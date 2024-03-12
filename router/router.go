package router

import (
	"golang-marketplace-app/controllers"
	"golang-marketplace-app/database"
	middleware "golang-marketplace-app/middlewere"
	"log"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	router := gin.Default()

	db, err := database.InitDB("postgres://postgres:P4ssW0rd@localhost:5434/marketplace_db?sslmode=disable")
	if err != nil {
		log.Fatal("Error initializing database:", err)
	}

	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	userRouter := router.Group("/v1/bank/account")
	{
		userRouter.POST("/", middleware.BankAccountValidator(), controllers.CreateBankAccount)
	}

	router.GET("/health-check", controllers.ServerCheck)

	return router
}
