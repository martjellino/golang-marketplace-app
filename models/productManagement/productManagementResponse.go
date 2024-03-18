package productmanage

import "time"

type ProductManagementResponse struct {
	ProductID      string    `json:"product_id"`
	SellerID       string    `json:"seller_id,omitempty"`
	Name           string    `json:"name"`
	Price          int       `json:"price"`
	ImageUrl       string    `json:"image_url"`
	Stock          int       `json:"stock"`
	Condition      string    `json:"condition"`
	Tags		   []string  `json:"tags"`
	IsPurchasable  bool      `json:"is_purchaseable"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type ProductStockManagementResponse struct {
	Stock          int       `json:"stock"`
	UpdatedAt      time.Time `json:"updated_at"`
}
