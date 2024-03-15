package services

import (
	"fmt"
	"golang-marketplace-app/database"
	bankaccount "golang-marketplace-app/models/productManagement"  //TODO: update variable name
	"log"
	"strconv"
	"time"
)

func FindProductByProductId(productID int) (bankaccount.ProductManagementResponse, error) {
	var (
			parsedProductId int
			sellerID        int
			name     		string
			price   		int
			imageUrl        string
			stock     		int
			condition   	string
			isPurchaseable  bool
			createdAt       time.Time
			updatedAt       time.Time
	)

	query := fmt.Sprintf("SELECT product_id, seller_id, name, price, image_url, stock, condition, is_purchaseable, created_at, updated_at FROM products WHERE product_id = %d", productID)
	fmt.Println("Query:", query) //

	err := database.DB.QueryRow(query).Scan(&parsedProductId, &sellerID, &name, &price, &imageUrl, &stock, &condition, &isPurchaseable, &createdAt, &updatedAt)
	if err != nil {
			log.Println(err)
			return bankaccount.ProductManagementResponse{}, fmt.Errorf("error retrieving product details: %v", err) //
	}

	return bankaccount.ProductManagementResponse{
		ProductID:     		strconv.Itoa(parsedProductId),
		SellerID:      		strconv.Itoa(sellerID),
		Name:		   		name,
		Price:				price,
		ImageUrl:       	imageUrl,
		Stock:		   		stock,
		Condition:			condition,
		IsPurchaseable:     isPurchaseable,
		CreatedAt:     		createdAt,
		UpdatedAt:     		updatedAt,
	}, nil
}

func DeleteProductByProductId(productID int) error {
	query := "DELETE FROM products WHERE product_id = $1"

	_, err := database.DB.Exec(query, productID)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("error deleting product: %v", err) //
	}

	return nil
}
