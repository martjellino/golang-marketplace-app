package controllers

import (
	"database/sql"
	"fmt"
	"golang-marketplace-app/database"
	"golang-marketplace-app/helpers"
	"golang-marketplace-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func UserRegister(ctx *gin.Context) {

	requestInterface, ok := ctx.Get("request")
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Parsed data not found in context",
		})
		return
	}

	user, ok := requestInterface.(models.Users)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to cast request connection to *bankaccount.BankAccountRequest",
		})
		return
	}

	// Check if username already exists
	if isUserExists(user.Username) {
		ctx.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
		return
	}

	// Validate password and username length
	if len(user.Password) < 5 || len(user.Password) > 15 || len(user.Username) < 5 || len(user.Username) > 15 || len(user.Fullname) < 5 || len(user.Fullname) > 15 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Password/Username/Name length should be between 5 and 15 characters"})
		return
	}

	// Before creating the user, perform necessary operations
	models.BeforeCreateUser(&user)

	// Save the user to the database
	if err := SaveUserToDatabase(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	// Generate JWT token
	token, err := helpers.GenerateToken(user.UserID, user.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": "Failed to generate token",
		})
		return
	}

	// Construct response data
	responseData := gin.H{
		"message": "User registered successfully",
		"data": gin.H{
			"username": user.Username,
			"name":     user.Fullname,
			"accessToken": token,
		},
	}
	ctx.JSON(http.StatusCreated, responseData)
}

func UserLogin(ctx *gin.Context) {
	var user models.Users
	db := database.GetDB()

	requestInterface, ok := ctx.Get("request")
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Parsed data not found in context",
		})
		return
	}

	Request, ok := requestInterface.(models.UserRequest)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to cast request connection to *bankaccount.BankAccountRequest",
		})
		return
	}

	// Retrieve user from the database based on the provided username
	err := db.QueryRow("SELECT user_id, username, password, fullname FROM users WHERE username = $1", Request.Username).
		Scan(&user.UserID, &user.Username, &user.Password, &user.Fullname)
	if err == sql.ErrNoRows {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": "User not found",
		})
		return
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": "Server error occurred",
		})
		return
	}

	// Compare password
	comparePass := helpers.ComparePassword([]byte(user.Password), []byte(Request.Password))
	if !comparePass {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Invalid password",
		})
		return
	}

	// Generate JWT token
	token, err := helpers.GenerateToken(user.UserID, user.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": "Failed to generate token",
		})
		return
	}

	// Construct response data
	responseData := gin.H{
		"message": "User logged successfully",
		"data": gin.H{
			"username":    user.Username,
			"name":        user.Fullname,
			"accessToken": token,
		},
	}

	ctx.JSON(http.StatusOK, responseData)
}

// SaveUserToDatabase saves the user data to the database
func SaveUserToDatabase(user *models.Users) error {
	db := database.GetDB()
	_, err := db.Exec("INSERT INTO users (username, password, fullname, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)", user.Username, user.Password, user.Fullname, user.CreatedAt, user.UpdatedAt)
	fmt.Println(user.Username)
	fmt.Println(user.Password)
	fmt.Println(user.Fullname)

	if err != nil {
		fmt.Println("Error inserting user into database:", err)
		return err
	}
	return nil
}

func isUserExists(username string) bool {
	db := database.GetDB()
	// Prepare SQL query to check if the username exists
	query := "SELECT EXISTS (SELECT 1 FROM users WHERE username = $1)"

	// Execute the query
	var exists bool
	err := db.QueryRow(query, username).Scan(&exists)
	if err != nil {
		// Handle the error, log it, etc.
		// For simplicity, let's assume the user does not exist in case of an error
		return false
	}

	return exists
}
