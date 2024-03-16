package controllers

import (
	bankaccount "golang-marketplace-app/models/productManagement"
	"golang-marketplace-app/services"
	"net/http"
	"strconv"
	"log"
	jwt5 "github.com/golang-jwt/jwt/v5"

	"github.com/gin-gonic/gin"
)

func CreateProduct(context *gin.Context) {
	requestInterface, ok := context.Get("request")
	if !ok {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Parsed data not found in context",
		})
		return
	}

	Request, ok := requestInterface.(bankaccount.ProductManagementRequest)
	if !ok {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to cast request to *bankaccount.ProductManagementRequest",
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

	var CreatedProduct, err = services.CreateProduct(userID, Request)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create product",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "product added successfully",
		"data": CreatedProduct,
	})
}

func DeleteProductByProductId(context *gin.Context) {
	productIdParam := context.Param("productId")
	productId, parseError := strconv.Atoi(productIdParam)

	if parseError != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to cast productId to int",
		})
		return
	}

	var ExistingProduct, findError = services.FindProductByProductId(productId)
	if findError != nil {
    if ExistingProduct.ProductID == "" {
			context.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch Product"})
		}
		return
	}

	var deleteError = services.DeleteProductByProductId(productId)
	if deleteError != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to delete Product",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "product deleted successfully"})
}
