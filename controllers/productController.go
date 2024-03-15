package controllers

import (
	"database/sql"
	"golang-marketplace-app/database"
	"golang-marketplace-app/models"
	"golang-marketplace-app/models/bankAccount"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetProducts(ctx *gin.Context) {
	db := database.GetDB()
	pageSize := 10
	pageNum, _ := strconv.Atoi(ctx.Param("pageNum"))
	offset := (pageNum - 1) * pageSize

	if offset < 0 {
		offset = 0
	}

	rows, err := db.Query("SELECT product_id, name, price, image_url, stock, condition, is_purchaseable, created_at, updated_at FROM products LIMIT $1 OFFSET $2", pageSize, offset)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}
	defer rows.Close()

	products := []models.Products{}
	for rows.Next() {
		var product models.Products
		err := rows.Scan(
			&product.ProductID,
			// &product.SellerID,
			&product.Name,
			&product.Price,
			&product.ImageUrl,
			&product.Stock,
			&product.Condition,
			// &product.Tags,
			&product.IsPurchaseable,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Internal Server Error",
				"message": err.Error(),
			})
			return
		}

		// Retrieve tags for the current product
		tags, err := getTagsForProduct(product.ProductID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Internal Server Error",
				"message": err.Error(),
			})
			return
		}

		// Assign tags to the product
		product.Tags = tags

		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	var totalCount int
	err = db.QueryRow("SELECT COUNT(*) FROM products").Scan(&totalCount)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    products,
		"meta": gin.H{
			"limit":  pageSize,
			"offset": offset,
			"total":  totalCount,
		},
	})
}

// Function to retrieve tags for a specific product
func getTagsForProduct(productID int) ([]models.Tags, error) {
	db := database.GetDB()
	rows, err := db.Query("SELECT tags.tag_id, tags.tag_name, tags.created_at, tags.updated_at FROM product_tags JOIN tags ON product_tags.tag_id = tags.tag_id WHERE product_tags.product_id = $1", productID)
	// rows, err := db.Query("SELECT tags.tag_name FROM product_tags JOIN tags ON product_tags.tag_id = tags.tag_id WHERE product_tags.product_id = $1", productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tags []models.Tags
	for rows.Next() {
		var tag models.Tags
		// err := rows.Scan(
		// 	&tag.Name,
		// )
		err := rows.Scan(
			&tag.TagID,
			&tag.Name,
			&tag.CreatedAt,
			&tag.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tags, nil
}

func GetProductByID(ctx *gin.Context) {
	db := database.GetDB()

	// Retrieve product ID from request parameters
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Invalid product ID",
		})
		return
	}

	// Query the database to fetch the product details
	var product models.Products
	err = db.QueryRow("SELECT product_id, seller_id, name, price, image_url, stock, condition, is_purchaseable, created_at, updated_at FROM products WHERE product_id = $1", id).
		Scan(
			&product.ProductID,
			&product.SellerID,
			&product.Name,
			&product.Price,
			&product.ImageUrl,
			&product.Stock,
			&product.Condition,
			&product.IsPurchaseable,
			&product.CreatedAt,
			&product.UpdatedAt,
		)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error":   "Not Found",
				"message": "Product not found",
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	// Retrieve tags for the current product
	tags, err := getTagsForProduct(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	// Assign tags to the product
	var tagNames []string
	for _, tag := range tags {
		tagNames = append(tagNames, tag.Name)
	}

	// Retrieve seller details from the users table
	var sellerName string
	err = db.QueryRow("SELECT fullname FROM users WHERE user_id = $1", product.SellerID).Scan(&sellerName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	// Retrieve bank accounts for the seller
	rows, err := db.Query("SELECT account_id, bank_name, account_name, account_number FROM bank_accounts WHERE user_id = $1", product.SellerID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}
	defer rows.Close()

	var bankAccounts []gin.H
	for rows.Next() {
		var bankAccount bankaccount.BankAccountResponse
		err := rows.Scan(&bankAccount.AccountID, &bankAccount.BankName, &bankAccount.AccountName, &bankAccount.AccountNumber)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Internal Server Error",
				"message": err.Error(),
			})
			return
		}
		bankAccounts = append(bankAccounts, gin.H{
			"bankAccountId":     bankAccount.AccountID,
			"bankName":          bankAccount.BankName,
			"bankAccountName":   bankAccount.AccountName,
			"bankAccountNumber": bankAccount.AccountNumber,
		})
	}
	if err := rows.Err(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data": gin.H{
			"product": gin.H{
				"productId":     product.ProductID,
				"name":          product.Name,
				"price":         product.Price,
				"imageUrl":      product.ImageUrl,
				"stock":         product.Stock,
				"condition":     product.Condition,
				"tags":          tagNames,
				"isPurchasable": product.IsPurchaseable,
			},
			"seller": gin.H{
				"name":         sellerName,
				"bankAccounts": bankAccounts,
			},
		},
	})
}
