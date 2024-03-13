package router

import (
	"database/sql"
	"golang-marketplace-app/controllers"
	middleware "golang-marketplace-app/middleware"

	"github.com/gin-gonic/gin"
)

func StartApp(db *sql.DB) *gin.Engine {
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	bankAccountRouter := router.Group("/v1/bank/account")
	{
		bankAccountRouter.POST("/", middleware.BankAccountValidator(), controllers.CreateBankAccount)
		bankAccountRouter.GET("/", controllers.GetBankAccountByUserId)
		bankAccountRouter.PATCH("/:accountId", middleware.BankAccountValidator(), controllers.UpdateBankAccountByAccountId)
		bankAccountRouter.DELETE("/:accountId", controllers.DeleteBankAccountByAccountId)
	}

	router.GET("/health-check", controllers.ServerCheck)

	return router
}
