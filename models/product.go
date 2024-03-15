package models

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type Products struct {
	ProductID      int       `json:"productId"`
	SellerID       int       `json:"sellerId"`
	Name           string    `json:"name" validate:"required,min=5,max=60"`
	Price          int       `json:"price" validate:"required,min=0"`
	ImageUrl       string    `json:"imageUrl" validate:"required,url"`
	Stock          int       `json:"stock" validate:"required,min=0"`
	Condition      string    `json:"condition" validate:"required,oneof=new second"`
	Tags           []Tags    `json:"tags"`
	IsPurchaseable bool      `json:"isPurchaseable" validate:"required"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func ValidateProduct(product *Products) error {
	validate := validator.New()
	return validate.Struct(product)
}
