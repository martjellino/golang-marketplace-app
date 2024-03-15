package controllers

import (
	"golang-marketplace-app/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
