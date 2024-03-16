package middleware

import (
	"golang-marketplace-app/helpers"
	product "golang-marketplace-app/models/productManagement"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProductValidator() gin.HandlerFunc {
	return func(context *gin.Context) {
		var productRequest product.ProductManagementRequest

		if payloadValidationError := context.ShouldBindJSON(&productRequest); payloadValidationError != nil {
			var errors []string

			if payloadValidationError.Error() == "EOF" {
				errors = append(errors, "Request body is missing")
			} else {
				errors = helpers.GeneralValidator(payloadValidationError)
			}

			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   errors,
				"message": "Failed to validate",
			})
			return
		}

		context.Set("request", productRequest)
		context.Next()
	}
}

func ProductUpdateValidator() gin.HandlerFunc {
	return func(context *gin.Context) {
		var productRequest product.ProductUpdateManagementRequest

		if payloadValidationError := context.ShouldBindJSON(&productRequest); payloadValidationError != nil {
			var errors []string

			if payloadValidationError.Error() == "EOF" {
				errors = append(errors, "Request body is missing")
			} else {
				errors = helpers.GeneralValidator(payloadValidationError)
			}

			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   errors,
				"message": "Failed to validate",
			})
			return
		}

		context.Set("request", productRequest)
		context.Next()
	}
}

func ProductStockValidator() gin.HandlerFunc {
	return func(context *gin.Context) {
		var productRequest product.ProductStockUpdateRequest

		if payloadValidationError := context.ShouldBindJSON(&productRequest); payloadValidationError != nil {
			var errors []string

			if payloadValidationError.Error() == "EOF" {
				errors = append(errors, "Request body is missing")
			} else {
				errors = helpers.GeneralValidator(payloadValidationError)
			}

			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   errors,
				"message": "Failed to validate",
			})
			return
		}

		context.Set("request", productRequest)
		context.Next()
	}
}
