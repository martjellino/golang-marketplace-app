package bankaccount

import "time"

type ProductManagementResponse struct {
	ProductID      string    `json:"product_id"`
	SellerID       string    `json:"seller_id,omitempty"`
	Name           string    `json:"name"`
	Price          int       `json:"price"`
	ImageUrl       string    `json:"image_url"`
	Stock          int       `json:"stock"`
	Condition      string    `json:"condition"`
	IsPurchasable  bool      `json:"is_purchaseable"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
