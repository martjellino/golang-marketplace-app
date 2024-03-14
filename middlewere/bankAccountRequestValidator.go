package middleware

import (
	"golang-marketplace-app/helpers"
	"golang-marketplace-app/models/bankAccount"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BankAccountValidator() gin.HandlerFunc {
	return func(context *gin.Context) {
		var bankAccountRequest bankaccount.BankAccountRequest

		if payloadValidationError := context.ShouldBindJSON(&bankAccountRequest); payloadValidationError != nil {
			var errors []string

			if payloadValidationError.Error() == "EOF" {
				errors = append(errors, "Request body is missing")
			} else {
				errors = helpers.GeneralValidator(payloadValidationError)
			}

			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   errors,
				"message": "Data Not Found",
			})
			return
		}

		context.Set("request", bankAccountRequest)
		context.Next()
	}
}
