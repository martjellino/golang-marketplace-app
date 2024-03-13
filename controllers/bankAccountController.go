package controllers

import (
	"database/sql"
	bankaccount "golang-marketplace-app/models/bankAccount"
	"golang-marketplace-app/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateBankAccount(context *gin.Context) {
	dbInterface, ok := context.Get("db")
	if !ok {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Database connection not found",
		})
		return
	}

	db, ok := dbInterface.(*sql.DB)
	if !ok {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to cast database connection to *sql.DB",
		})
		return
	}

	requestInterface, ok := context.Get("request")
	if !ok {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Parsed data not found in context",
		})
		return
	}

	Request, ok := requestInterface.(bankaccount.BankAccountRequest)
	if !ok {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to cast request connection to *bankaccount.BankAccountRequest",
		})
		return
	}

	var CreatedBankAccount, err = services.CreateBankAccount(Request, db)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create bank account",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "account added successfully",
		"data": CreatedBankAccount,
	})
}

func UpdateBankAccountByAccountId(context *gin.Context) {
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
			"message": "Parsed data not found in context",
		})
		return
	}

	Request, ok := requestInterface.(bankaccount.BankAccountRequest)
	if !ok {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to cast request connection to *bankaccount.BankAccountRequest",
		})
		return
	}

	accountIdParam := context.Param("accountId")
	accountId, parseError := strconv.Atoi(accountIdParam)

	if parseError != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to cast accountId to int",
		})
		return
	}

	var ExistingBankAccount, findError = services.FindBankAccountByAccountId(accountId, db)
	if findError != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to fetch bank account",
		})
		return
	}

	if ExistingBankAccount.AccountID == "" {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Bank account not found",
		})
		return
	}

	var UpdatedBankAccount, updateError = services.UpdateBankAccountByAccountId(accountId, Request, db)
	if updateError != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to update bank account",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "account updated successfully",
		"data": UpdatedBankAccount,
	})
}