package router

import (
	"golang-marketplace-app/controllers"
	middleware "golang-marketplace-app/middleware"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requestHistogram = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "http_request_duration_seconds",
		Help:    "Histogram of request duration.",
		Buckets: prometheus.LinearBuckets(1, 1, 10),
	}, []string{"path", "method", "status"})
)

func StartApp() *gin.Engine {
	router := gin.Default()

	userRouter := router.Group("v1/user")
	{
		userRouter.POST("/register", wrapHandlerWithMetrics("v1/user/register", "POST", controllers.UserRegister))
		userRouter.POST("/login", wrapHandlerWithMetrics("v1/user/login", "POST", controllers.UserLogin))
	}

	bankAccountRouter := router.Group("/v1/bank/account")
	{
		bankAccountRouter.POST("/", middleware.Authentication(), middleware.BankAccountValidator(), wrapHandlerWithMetrics("/v1/bank/account", "POST", controllers.CreateBankAccount))
		bankAccountRouter.GET("/", middleware.Authentication(), wrapHandlerWithMetrics("/v1/bank/account", "GET", controllers.GetBankAccountByUserId))
		bankAccountRouter.PATCH("/:accountId", middleware.Authentication(), middleware.BankAccountValidator(), wrapHandlerWithMetrics("v1/bank/account/:acccountId", "PATCH", controllers.UpdateBankAccountByAccountId))
		bankAccountRouter.DELETE("/:accountId", middleware.Authentication(), wrapHandlerWithMetrics("v1/bank/account/:accountId", "DELETE", controllers.DeleteBankAccountByAccountId))
	}

	productManagementRouter := router.Group("v1/product")
	{
		// productManagementRouter.POST("/", controllers.CreateProduct)
		// productManagementRouter.PATCH("/:productId", controllers.UpdateProductByProductId)
		productManagementRouter.DELETE("/:productId", controllers.DeleteProductByProductId) //TODO: not implement middleware  yet

		productManagementRouter.GET("/", controllers.ListProduct)
		productManagementRouter.GET("/:productId", controllers.DetailProductByProductId)
	}

	paymentRouter := router.Group("/v1/product")
	{
		paymentRouter.POST("/:productId/buy", middleware.Authentication(), middleware.PaymentValidator(), controllers.CreatePaymentToAProductId)
	}

	router.GET("/health-check", controllers.ServerCheck)
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	return router
}

func wrapHandlerWithMetrics(path, method string, handler gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		handler(c)
		duration := time.Since(startTime).Seconds()
		log.Println(path, method, strconv.Itoa(c.Writer.Status()))
		requestHistogram.WithLabelValues(path, method, strconv.Itoa(c.Writer.Status())).Observe(duration)
	}
}
