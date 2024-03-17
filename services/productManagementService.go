package services

import (
	"fmt"
	"golang-marketplace-app/database"
	productmanage "golang-marketplace-app/models/productManagement"
	"log"
	"strconv"
	"time"
)

func CreateProduct(userId int, Request productmanage.ProductManagementRequest) (productmanage.ProductManagementResponse, error) {
	var productID int
	productSequenceError := database.DB.QueryRow("SELECT nextval('products_product_id_seq')").Scan(&productID)
	if productSequenceError != nil {
		log.Println("Error retrieving last inserted ID products:", productSequenceError)
		return productmanage.ProductManagementResponse{}, fmt.Errorf("error retrieving last inserted ID products: %v", productSequenceError)
	}

	stmt, err := database.DB.Prepare("INSERT INTO products (seller_id, name, price, image_url, stock, condition, is_purchaseable) VALUES ($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
		log.Println("Error preparing SQL query:", err)
		return productmanage.ProductManagementResponse{}, fmt.Errorf("error preparing SQL query: %v", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(userId, Request.Name, Request.Price, Request.ImageUrl, Request.Stock, Request.Condition, Request.IsPurchaseable)
	if err != nil {
		log.Println("Error executing insert statement:", err)
		return productmanage.ProductManagementResponse{}, fmt.Errorf("error executing insert statement: %v", err)
	}

	parsedProductID := strconv.Itoa(productID)
	parsedUserId := strconv.Itoa(userId)

	for _, tag := range Request.Tags {
		var tagID int
		err = database.DB.QueryRow("SELECT nextval('tags_tag_id_seq')").Scan(&tagID)
		if productSequenceError != nil {
			log.Println("Error retrieving last inserted ID tags:", err)
			return productmanage.ProductManagementResponse{}, fmt.Errorf("error retrieving last inserted ID tags: %v", err)
		}

		// bulk insert
		stmt2, err := database.DB.Prepare("INSERT INTO tags (tag_name) VALUES ($1)")
		if err != nil {
			log.Println("Error preparing SQL query:", err)
			return productmanage.ProductManagementResponse{}, fmt.Errorf("error preparing SQL query: %v", err)
		}
		defer stmt2.Close()
		_, err = stmt2.Exec(tag)
		if err != nil {
			log.Println("Error executing insert statement:", err)
			return productmanage.ProductManagementResponse{}, fmt.Errorf("error executing insert statement: %v", err)
		}
		// parsedTagID := strconv.Itoa(tagID)

		stmt3, err := database.DB.Prepare("INSERT INTO product_tags (product_id, tag_id) VALUES ($1, $2)")
		if err != nil {
			log.Println("Error preparing SQL query:", err)
			return productmanage.ProductManagementResponse{}, fmt.Errorf("error preparing SQL query: %v", err)
		}
		defer stmt3.Close()
		_, err = stmt3.Exec(productID, tagID)
		if err != nil {
			log.Println("Error executing insert statement:", err)
			return productmanage.ProductManagementResponse{}, fmt.Errorf("error executing insert statement: %v", err)
		}
	}

	return productmanage.ProductManagementResponse{
		ProductID:     parsedProductID,
		Name:          Request.Name,
		Price:         Request.Price,
		ImageUrl:      Request.ImageUrl,
		Stock:         Request.Stock,
		Condition:     Request.Condition,
		Tags:          Request.Tags,
		IsPurchasable: Request.IsPurchaseable,
		SellerID:      parsedUserId,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}, nil
}

func UpdateProductByProductId(productID int, Request productmanage.ProductUpdateManagementRequest) (productmanage.ProductManagementResponse, error) {
	stmt, err := database.DB.Prepare("UPDATE products SET name=$1, price=$2, image_url=$3, condition=$4, is_purchaseable=$5, updated_at=$6 WHERE product_id=$7")

	if err != nil {
		return productmanage.ProductManagementResponse{}, fmt.Errorf("error preparing SQL query: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(Request.Name, Request.Price, Request.ImageUrl, Request.Condition, Request.IsPurchaseable, time.Now(), productID)
	if err != nil {
		return productmanage.ProductManagementResponse{}, fmt.Errorf("error executing update product statement: %v", err)
	}

	return productmanage.ProductManagementResponse{
		ProductID:     strconv.Itoa(productID),
		Name:          Request.Name,
		Price:         Request.Price,
		ImageUrl:      Request.ImageUrl,
		Condition:     Request.Condition,
		Tags:          Request.Tags,
		IsPurchasable: Request.IsPurchaseable,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}, nil
}

func FindProductByProductId(productID int) (productmanage.ProductManagementResponse, error) {
	var (
		parsedProductId int
		sellerID        int
		name            string
		price           int
		imageUrl        string
		stock           int
		condition       string
		isPurchaseable  bool
		createdAt       time.Time
		updatedAt       time.Time
	)

	query := fmt.Sprintf("SELECT product_id, seller_id, name, price, image_url, stock, condition, is_purchaseable, created_at, updated_at FROM products WHERE product_id = %d", productID)

	err := database.DB.QueryRow(query).Scan(&parsedProductId, &sellerID, &name, &price, &imageUrl, &stock, &condition, &isPurchaseable, &createdAt, &updatedAt)
	if err != nil {
		log.Println(err)
		return productmanage.ProductManagementResponse{}, fmt.Errorf("error retrieving bank account details: %v", err)
	}

	return productmanage.ProductManagementResponse{
		ProductID:     strconv.Itoa(parsedProductId),
		SellerID:      strconv.Itoa(sellerID),
		Name:          name,
		Price:         price,
		ImageUrl:      imageUrl,
		Stock:         stock,
		Condition:     condition,
		IsPurchasable: isPurchaseable,
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
	}, nil
}

func DeleteProductByProductId(productID int) error {
	query := "DELETE FROM products WHERE product_id = $1"

	_, err := database.DB.Exec(query, productID)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("error deleting product: %v", err)
	}

	return nil
}

func UpdateStockProductByProductId(productID int, Request productmanage.ProductStockUpdateRequest) (productmanage.ProductStockManagementResponse, error) {
	stmt, err := database.DB.Prepare("UPDATE products SET stock=$1, updated_at=$2 WHERE product_id=$3")
	if err != nil {
		return productmanage.ProductStockManagementResponse{}, fmt.Errorf("error preparing SQL query: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(Request.Stock, time.Now(), productID)
	if err != nil {
		return productmanage.ProductStockManagementResponse{}, fmt.Errorf("error executing update stock statement: %v", err)
	}

	return productmanage.ProductStockManagementResponse{
		Stock:     Request.Stock,
		UpdatedAt: time.Now(),
	}, nil
}
