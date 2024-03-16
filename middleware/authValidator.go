package middleware

import (
	"golang-marketplace-app/helpers"
	"golang-marketplace-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthValidator() gin.HandlerFunc {
	return func(context *gin.Context) {
		var paymentRequest models.UserRequest

		if payloadValidationError := context.ShouldBindJSON(&paymentRequest); payloadValidationError != nil {
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

		context.Set("request", paymentRequest)
		context.Next()
	}
}
