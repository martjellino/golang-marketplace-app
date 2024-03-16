package controllers

import (
	bankaccount "golang-marketplace-app/models/bankAccount"
	"golang-marketplace-app/services"
	"log"
	"net/http"
	"strconv"
	jwt5 "github.com/golang-jwt/jwt/v5"

	"github.com/gin-gonic/gin"
)

func CreateBankAccount(context *gin.Context) {
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
			"message": "Failed to cast request to *bankaccount.BankAccountRequest",
		})
		return
	}

	JwtPayload, ok := context.Get("userData")
	log.Println(JwtPayload)
	if !ok {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Parsed user data not found in context",
		})
		return
	}

	userData := context.MustGet("userData").(jwt5.MapClaims)
	userID := int(userData["id"].(float64))

	var CreatedBankAccount, err = services.CreateBankAccount(userID, Request)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create bank account",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "account added successfully",
		"data":    CreatedBankAccount,
	})
}

func UpdateBankAccountByAccountId(context *gin.Context) {
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

	var ExistingBankAccount, findError = services.FindBankAccountByAccountId(accountId)
	if findError != nil {
		if ExistingBankAccount.AccountID == "" {
			context.JSON(http.StatusNotFound, gin.H{"message": "Bank account not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch bank account"})
		}
		return
	}

	userData := context.MustGet("userData").(jwt5.MapClaims)
	userID := int(userData["id"].(float64))

	if ExistingBankAccount.UserID != strconv.Itoa(userID) {
		context.JSON(http.StatusForbidden, gin.H{
			"message": "Forbidden",
		})
		return
	}

	var UpdatedBankAccount, updateError = services.UpdateBankAccountByAccountId(accountId, Request)
	if updateError != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to update bank account",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "account updated successfully",
		"data":    UpdatedBankAccount,
	})
}

func DeleteBankAccountByAccountId(context *gin.Context) {
	accountIdParam := context.Param("accountId")
	accountId, parseError := strconv.Atoi(accountIdParam)

	if parseError != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to cast accountId to int",
		})
		return
	}

	var ExistingBankAccount, findError = services.FindBankAccountByAccountId(accountId)
	if findError != nil {
    if ExistingBankAccount.AccountID == "" {
			context.JSON(http.StatusNotFound, gin.H{"message": "Bank account not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch bank account"})
		}
		return
	}

	userData := context.MustGet("userData").(jwt5.MapClaims)
	userID := int(userData["id"].(float64))

	if ExistingBankAccount.UserID != strconv.Itoa(userID) {
		context.JSON(http.StatusForbidden, gin.H{
			"message": "Forbidden",
		})
		return
	}

	var deleteError = services.DeleteBankAccountByAccountId(accountId)
	if deleteError != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to delete bank account",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "account deleted successfully"})
}

func GetBankAccountByUserId(context *gin.Context) {
	userData := context.MustGet("userData").(jwt5.MapClaims)
	userID := int(userData["id"].(float64))

	var bankAccounts, getBankAccounsError = services.GetBankAccountsByUserId(userID)
	if getBankAccounsError != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to fetch bank account",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":   bankAccounts ,
	})
}