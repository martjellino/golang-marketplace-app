package services

import (
	"fmt"
	"golang-marketplace-app/database"
	"golang-marketplace-app/models/payment"
	"log"
)

func CreatePayment(paymentResp payment.PaymentResponse) (payment.PaymentResponse, error) {
	
	stmt, err := database.DB.Prepare("INSERT INTO purchases (buyer_id, account_id, product_id, qty, total_price, image_url) VALUES ($1, $2, $3, $4, $5, $6) RETURNING purchase_id, created_at, updated_at")
	if err != nil {
		log.Println("Error preparing SQL query:", err)
		return paymentResp, fmt.Errorf("error preparing SQL query: %v", err)
	}
	defer stmt.Close()

	var newPaymentResp payment.PaymentResponse
	err = stmt.QueryRow(
		paymentResp.BuyerID,
		paymentResp.AccountID,
		paymentResp.ProductID,
		paymentResp.Quantity,
		paymentResp.TotalPrice,
		paymentResp.ImageUrl,
	).Scan(
		&newPaymentResp.PurchaseID,
		&newPaymentResp.CreatedAt,
		&newPaymentResp.UpdatedAt,
	)
	if err != nil {
		log.Println("Error executing insert statement:", err)
		return newPaymentResp, fmt.Errorf("error executing insert statement: %v", err)
	}

	newPaymentResp.AccountID = paymentResp.AccountID
	newPaymentResp.BuyerID = paymentResp.BuyerID
	newPaymentResp.ProductID = paymentResp.ProductID
	newPaymentResp.Quantity = paymentResp.Quantity
	newPaymentResp.TotalPrice = paymentResp.TotalPrice
	newPaymentResp.ImageUrl = paymentResp.ImageUrl

	return newPaymentResp, nil
}
