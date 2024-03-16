package controllers

import (
	"fmt"
	productmanage "golang-marketplace-app/models/productManagement"
	"golang-marketplace-app/services"
	"net/http"
	"net/url"
	"strconv"
	"strings"

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

	Request, ok := requestInterface.(productmanage.ProductManagementRequest)
	if !ok {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to cast request to *productmanage.ProductManagementRequest",
		})
		return
	}

	if !(Request.Condition == "new" || Request.Condition == "second") {
			context.JSON(http.StatusBadRequest, gin.H{
					"message": "Status should be new or second",
			})
			return
	}
	
	parsedURL, parsedUrlError := url.ParseRequestURI(Request.ImageUrl)
	if parsedUrlError != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
			context.JSON(http.StatusBadRequest, gin.H{
					"message": "ImageUrl should be a valid URL with scheme and host",
			})
			return
	}
	
	if parsedURL.Host != "s3.amazonaws.com" {
			context.JSON(http.StatusBadRequest, gin.H{
					"message": "ImageUrl should be hosted on Amazon S3",
			})
			return
	}
	
	if !strings.HasPrefix(parsedURL.Path, "/sprint-bucket-public-read/") {
			context.JSON(http.StatusBadRequest, gin.H{
					"message": "ImageUrl path should start with '/sprint-bucket-public-read/'",
			})
			return
	}

	userData := context.MustGet("userData").(jwt5.MapClaims)
	userID := int(userData["id"].(float64))

	var CreatedProduct, err = services.CreateProduct(userID, Request)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Failed to create product %s", err),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "product added successfully",
		"data":    CreatedProduct,
	})
}

func UpdateProductByProductId(context *gin.Context) {
	requestInterface, ok := context.Get("request")

	if !ok {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Parsed data not found in context",
		})
		return
	}

	Request, ok := requestInterface.(productmanage.ProductUpdateManagementRequest)
	if !ok {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to cast request connection to *productmanage.ProductUpdateManagementRequest",
		})
		return
	}

	if !(Request.Condition == "new" || Request.Condition == "second") {
			context.JSON(http.StatusBadRequest, gin.H{
					"message": "Status should be new or second",
			})
			return
	}
	
	parsedURL, parsedUrlError := url.ParseRequestURI(Request.ImageUrl)
	if parsedUrlError != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
			context.JSON(http.StatusBadRequest, gin.H{
					"message": "ImageUrl should be a valid URL with scheme and host",
			})
			return
	}
	
	if parsedURL.Host != "s3.amazonaws.com" {
			context.JSON(http.StatusBadRequest, gin.H{
					"message": "ImageUrl should be hosted on Amazon S3",
			})
			return
	}
	
	if !strings.HasPrefix(parsedURL.Path, "/sprint-bucket-public-read/") {
			context.JSON(http.StatusBadRequest, gin.H{
					"message": "ImageUrl path should start with '/sprint-bucket-public-read/'",
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

	var ExistingProduct, findError = services.FindProductByProductId(productId)
	if findError != nil {
		if ExistingProduct.ProductID == "" {
			context.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch product"})
		}
		return
	}

	userData := context.MustGet("userData").(jwt5.MapClaims)
	userID := int(userData["id"].(float64))

	if ExistingProduct.SellerID != strconv.Itoa(userID) {
		context.JSON(http.StatusForbidden, gin.H{
			"message": "Forbidden",
		})
		return
	}

	var UpdatedProduct, updateError = services.UpdateProductByProductId(productId, Request)
	if updateError != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to update stock product",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "stock updated successfully",
		"data":    UpdatedProduct,
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

	userData := context.MustGet("userData").(jwt5.MapClaims)
	userID := int(userData["id"].(float64))

	if ExistingProduct.SellerID != strconv.Itoa(userID) {
		context.JSON(http.StatusForbidden, gin.H{
			"message": "Forbidden",
		})
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

func UpdateStockProductByProductId(context *gin.Context) {
	requestInterface, ok := context.Get("request")

	if !ok {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Parsed data not found in context",
		})
		return
	}

	Request, ok := requestInterface.(productmanage.ProductStockUpdateRequest)
	if !ok {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to cast request connection to *productmanage.ProductStockUpdateRequest",
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

	var ExistingProduct, findError = services.FindProductByProductId(productId)
	if findError != nil {
		if ExistingProduct.ProductID == "" {
			context.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch product"})
		}
		return
	}

	userData := context.MustGet("userData").(jwt5.MapClaims)
	userID := int(userData["id"].(float64))

	if ExistingProduct.SellerID != strconv.Itoa(userID) {
		context.JSON(http.StatusForbidden, gin.H{
			"message": "Forbidden",
		})
		return
	}

	var UpdatedProduct, updateError = services.UpdateStockProductByProductId(productId, Request)
	if updateError != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to update stock product",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "stock updated successfully",
		"data":    UpdatedProduct,
	})
}
