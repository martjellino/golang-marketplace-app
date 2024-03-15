package controllers

import (
	"fmt"
	"golang-marketplace-app/models/payment"
	"golang-marketplace-app/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	jwt5 "github.com/golang-jwt/jwt/v5"
)

func CreatePaymentToAProductId(context *gin.Context) {
	requestInterface, ok := context.Get("request")
	if !ok {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Parsed data not found in context",
		})
		return
	}

	Request, ok := requestInterface.(payment.PaymentRequest)
	if !ok {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to cast request to *bankaccount.BankAccountRequest",
		})
		return
	}

	productIdParam := context.Param("productId")
	productId, parseError := strconv.Atoi(productIdParam)

	if parseError != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to cast productId to int",
		})
		return
	}

	userData := context.MustGet("userData").(jwt5.MapClaims)
	userID := int(userData["id"].(float64))

	ExistingProduct, err := services.GetProductById(productId)
	if err != nil {
		if ExistingProduct.Name == "" {
			context.JSON(http.StatusNotFound, gin.H{
				"message": fmt.Sprintf("Product with ID %d not found", productId),
			})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch bank account"})
		}
		return
	}

	accoundIdString := Request.BankAccountId
	accoundId, parseError := strconv.Atoi(accoundIdString)

	if parseError != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to cast accountId to int",
		})
		return
	}

	var ExistingBankAccount, findError = services.FindBankAccountByAccountId(accoundId)
	if findError != nil {
		if ExistingBankAccount.AccountID == "" {
			context.JSON(http.StatusNotFound, gin.H{"message": "Bank account not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch bank account"})
		}
		return
	}

	err = services.UpdateStock(productId, Request.Quantity)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	paymentResponse := payment.PaymentResponse {
		BuyerID: strconv.Itoa(userID),
		AccountID: strconv.Itoa(accoundId),
		ProductID: strconv.Itoa(ExistingProduct.ProductID),
		Quantity: Request.Quantity,
		TotalPrice: Request.Quantity * ExistingProduct.Price,
		ImageUrl: Request.PaymentProofImageUrl,
	}

	var PaymentResult, createPaymentError = services.CreatePayment(paymentResponse)
	if createPaymentError != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create payment account",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "account added successfully",
		"data":    PaymentResult,
	})
}
