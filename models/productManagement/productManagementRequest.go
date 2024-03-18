package productmanage

import "time"

type ProductManagementRequest struct {
	ProductId      int       `json:"productId"`
	SellerID       int       `json:"sellerID"`
	Name           string    `json:"name" binding:"required,min=5,max=60" validate:"required,min=5,max=60"`
	Price          int       `json:"price" binding:"required,min=0"`
	ImageUrl       string    `json:"imageUrl" binding:"required,min=5,max=255" validate:"required,min=5,max=255"`
	Stock          int       `json:"stock" binding:"required,min=0"`
	Condition      string    `json:"condition" binding:"required" validate:"required"`
	Tags           []string  `json:"tags" binding:"required"`
	IsPurchaseable bool      `json:"isPurchaseable" binding:"required"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

type ProductUpdateManagementRequest struct {
	ProductId      int       `json:"productId"`
	SellerID       int       `json:"sellerID"`
	Name           string    `json:"name" binding:"required,min=5,max=60" validate:"required,min=5,max=60"`
	Price          int       `json:"price" binding:"required,min=0"`
	ImageUrl       string    `json:"imageUrl" binding:"required,min=5" validate:"required,min=5"`
	Condition      string    `json:"condition" binding:"required" validate:"required"`
	Tags           []string  `json:"tags" binding:"required"`
	IsPurchaseable bool      `json:"isPurchaseable" binding:"required"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

type ProductStockUpdateRequest struct {
	Stock int `json:"stock" binding:"required,min=0"`
}
