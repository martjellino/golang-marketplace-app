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

	productManagementRouter := router.Group("v1/product")
	{
		// productManagementRouter.POST("/", controllers.CreateProduct)
		// productManagementRouter.PATCH("/:productId", controllers.UpdateProductByProductId)
		productManagementRouter.DELETE("/:productId", controllers.DeleteProductByProductId) //TODO: not implement middleware  yet

		// productManagementRouter.GET("/", controllers.ListProduct)
		// productManagementRouter.GET("/:productId", controllers.DetailProductByProductId)
	}

	router.GET("/health-check", controllers.ServerCheck)

	return router
}
