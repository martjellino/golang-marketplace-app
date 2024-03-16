package middleware

import (
	"golang-marketplace-app/helpers"
	"golang-marketplace-app/models/bankAccount"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func BankAccountPatchValidator() gin.HandlerFunc {
	return func(context *gin.Context) {
		var bankAccountRequest bankaccount.BankAccountRequest

		if payloadValidationError := context.ShouldBindJSON(&bankAccountRequest); payloadValidationError != nil {
			var errors []string
			
			accountIdParam := context.Param("accountId")
			_, parseError := strconv.Atoi(accountIdParam)
			if parseError != nil {
				errors = append(errors, "Path is invalid")
				context.AbortWithStatusJSON(http.StatusNotFound, gin.H{
					"error":   errors,
					"message": "Failed to validate",
				})
				return
			}  else if payloadValidationError.Error() == "EOF" {
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

		context.Set("request", bankAccountRequest)
		context.Next()
	}
}
