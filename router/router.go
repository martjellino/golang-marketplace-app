package router

import (
	"golang-marketplace-app/controllers"
	middleware "golang-marketplace-app/middlewere"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	router := gin.Default()

	userRouter := router.Group("/v1/bank/account")
	{
		userRouter.POST("/", middleware.BankAccountValidator(), controllers.CreateBankAccount)
	}

	router.GET("/health-check", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "success",
		})
	})

	return router
}