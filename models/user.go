package models

import (
	"golang-marketplace-app/helpers"
	"time"
)

type Users struct {
	UserID    int       `json:"user_id"`
	Username  string    `json:"username" binding:"min=5,max=15" validate:"min=5,max=15"`
	Password  string    `json:"password" binding:"min=5,max=15" validate:"min=5,max=15"`
	Fullname  string    `json:"fullname,omitempty" binding:"omitempty,min=5,max=50" validate:"omitempty,min=5,max=50"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// HashPassword hashes the password before creating the user
func (u *Users) HashPassword() {
	// Hash the password using a hashing function like bcrypt
	u.Password = helpers.HashPassword(u.Password)
}

// BeforeCreateUser is a function to be called before creating a new user
func BeforeCreateUser(user *Users) {
    // Perform any pre-create logic here
    user.CreatedAt = time.Now()
    user.UpdatedAt = time.Now()
    user.HashPassword()
}