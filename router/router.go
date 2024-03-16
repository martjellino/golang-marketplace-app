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
		bankAccountRouter.POST("/", middleware.Authentication(), middleware.BankAccountValidator(), controllers.CreateBankAccount)
		bankAccountRouter.GET("/", middleware.Authentication(), controllers.GetBankAccountByUserId)
		bankAccountRouter.PATCH("/:accountId", middleware.Authentication(), middleware.BankAccountValidator(), controllers.UpdateBankAccountByAccountId)
		bankAccountRouter.DELETE("/:accountId", middleware.Authentication(), controllers.DeleteBankAccountByAccountId)
	}
	productManagementRouter := router.Group("v1/product")
	{	
		productManagementRouter.GET("/", controllers.ListProduct)
		productManagementRouter.GET("/:productId", controllers.DetailProductByProductId)
		productManagementRouter.POST("/", middleware.Authentication(), middleware.ProductValidator(), controllers.CreateProduct)
		productManagementRouter.PATCH("/:productId", middleware.Authentication(), middleware.ProductUpdateValidator(), controllers.UpdateProductByProductId)
		productManagementRouter.DELETE("/:productId", middleware.Authentication(), controllers.DeleteProductByProductId)
		productManagementRouter.POST("/:productId/stock", middleware.Authentication(), middleware.ProductStockValidator(), controllers.UpdateStockProductByProductId)
		productManagementRouter.POST("/:productId/buy", middleware.Authentication(), middleware.PaymentValidator(), controllers.CreatePaymentToAProductId)
	}

	imageUploadRouter := router.Group("/v1/image")
	{
		imageUploadRouter.POST("/", middleware.Authentication(), controllers.CreateUploadImage)
	}

	router.GET("/health-check", controllers.ServerCheck)

	return router
}
