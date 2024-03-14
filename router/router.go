package router

import (
	"golang-marketplace-app/controllers"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	router := gin.Default()

	userRouter := router.Group("v1/user")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}

	router.GET("/health-check", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "success",
		})
	})

	return router
}
