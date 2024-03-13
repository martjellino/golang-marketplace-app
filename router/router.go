package router

import (
	"database/sql"
	"golang-marketplace-app/controllers"
	middleware "golang-marketplace-app/middlewere"

	"github.com/gin-gonic/gin"
)

func StartApp(db *sql.DB) *gin.Engine {
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	userRouter := router.Group("/v1/bank/account")
	{
		userRouter.POST("/", middleware.BankAccountValidator(), controllers.CreateBankAccount)
		userRouter.GET("/", controllers.GetBankAccountByUserId)
		userRouter.PATCH("/:accountId", middleware.BankAccountValidator(), controllers.UpdateBankAccountByAccountId)
		userRouter.DELETE("/:accountId", controllers.DeleteBankAccountByAccountId)
	}

	router.GET("/health-check", controllers.ServerCheck)

	return router
}
