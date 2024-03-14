package router

import (
	"golang-marketplace-app/controllers"
	middleware "golang-marketplace-app/middleware"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	router := gin.Default()

	userRouter := router.Group("v1/user")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}

	bankAccountRouter := router.Group("/v1/bank/account")
	{
		bankAccountRouter.POST("/",  middleware.Authentication(), middleware.BankAccountValidator(), controllers.CreateBankAccount)
		bankAccountRouter.GET("/", middleware.Authentication(), controllers.GetBankAccountByUserId)
		bankAccountRouter.PATCH("/:accountId", middleware.Authentication(), middleware.BankAccountValidator(), controllers.UpdateBankAccountByAccountId)
		bankAccountRouter.DELETE("/:accountId", middleware.Authentication(), controllers.DeleteBankAccountByAccountId)
	}

	router.GET("/health-check", controllers.ServerCheck)

	return router
}
