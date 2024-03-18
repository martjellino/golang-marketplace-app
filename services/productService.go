package services

import (
	"database/sql"
	"fmt"
	"golang-marketplace-app/database"
	"golang-marketplace-app/models/product"
	"log"
)

func GetProductById(productId int) (product.Product, error) {
	var prod product.Product

	stmt, err := database.DB.Prepare("SELECT * FROM products WHERE product_id = $1")
	if err != nil {
		log.Println("Error preparing SQL query:", err)
		return prod, fmt.Errorf("error preparing SQL query: %v", err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(productId).Scan(
		&prod.ProductID,
		&prod.SellerID,
		&prod.Name,
		&prod.Price,
		&prod.ImageUrl,
		&prod.Stock,
		&prod.Condition,
		&prod.IsPurchaseable,
		&prod.CreatedAt,
		&prod.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return prod, fmt.Errorf("product with ID %d not found", productId)
		}
		log.Println("Error retrieving product:", err)
		return prod, fmt.Errorf("error retrieving product: %v", err)
	}

	return prod, nil
}

func UpdateStock(productId, decreaseBy int) error {
	stmt, err := database.DB.Prepare("UPDATE products SET stock = stock - $1 WHERE product_id = $2 AND stock >= $1")
	if err != nil {
		log.Println("Error preparing SQL query:", err)
		return fmt.Errorf("error preparing SQL query: %v", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(decreaseBy, productId)
	if err != nil {
		log.Println("Error updating stock:", err)
		return fmt.Errorf("error updating stock: %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Println("Error getting rows affected:", err)
		return fmt.Errorf("error getting rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("not enough stock available for product with ID %d", productId)
	}

	return nil
}
