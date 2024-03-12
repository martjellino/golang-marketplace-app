package controllers

import (
	"database/sql"
	bankaccount "golang-marketplace-app/models/bankAccount"
	"golang-marketplace-app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBankAccount(context *gin.Context) {
	dbInterface, ok := context.Get("db")
	if !ok {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Database connection not found",
		})
		return
	}

	db, ok := dbInterface.(*sql.DB)
	if !ok {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to cast database connection to *sql.DB",
		})
		return
	}

	requestInterface, ok := context.Get("request")
	if !ok {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Parsed data not found in context",
		})
		return
	}

	Request, ok := requestInterface.(bankaccount.BankAccountRequest)
	if !ok {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to cast request connection to *bankaccount.BankAccountRequest",
		})
		return
	}

	var Response, err = services.CreateBankAccount(Request, db)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create bank account",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "account added successfully",
		"data": Response,
	})
}
