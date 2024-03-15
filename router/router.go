package router

import (
	"database/sql"
	"golang-marketplace-app/controllers"
	middleware "golang-marketplace-app/middlewere"

	"github.com/gin-gonic/gin"
)

func StartApp(db *sql.DB) *gin.Engine {
	router := gin.Default()

	userRouter := router.Group("v1/user")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}
	bankAccountRouter := router.Group("/v1/bank/account")
	{
		bankAccountRouter.POST("/", middleware.BankAccountValidator(), controllers.CreateBankAccount)
	}
	productRouter := router.Group("v1/product")
	{
		productRouter.GET("/", controllers.GetProducts)
		productRouter.GET("/:id", controllers.GetProductByID)
	}

	router.GET("/health-check", controllers.ServerCheck)

	return router
}
