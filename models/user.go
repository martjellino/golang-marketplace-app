package models

import (
	"github.com/go-playground/validator/v10"
	"golang-marketplace-app/helpers"
	"time"
)

type Users struct {
	UserID    int       `json:"user_id"`
	Username  string    `json:"username" binding:"min=5,max=15" validate:"min=5,max=15"`
	Password  string    `json:"password" binding:"min=5,max=15" validate:"min=5,max=15"`
	Fullname  string    `json:"name" binding:"min=5,max=50" validate:"min=5,max=50"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRequest struct {
	Username  string    `json:"username" binding:"min=5,max=15" validate:"min=5,max=15"`
	Password  string    `json:"password" binding:"min=5,max=15" validate:"min=5,max=15"`
}

// HashPassword hashes the password before creating the user
func (u *Users) HashPassword() error {
	// Hash the password using a hashing function like bcrypt
	hashedPassword, err := helpers.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashedPassword
	return nil
}

// BeforeCreateUser is a function to be called before creating a new user
func BeforeCreateUser(user *Users) {
	// Perform any pre-create logic here
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.HashPassword()
}

func ValidateUser(user *Users) error {
	validate := validator.New()
	return validate.Struct(user)
}
