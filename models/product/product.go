
package product

import (
	"time"
)

type Product struct {
	ProductID      int
	SellerID       int
	Name           string
	Price          int
	ImageUrl       string
	Stock          int    
	Condition      string 
	IsPurchaseable bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
}