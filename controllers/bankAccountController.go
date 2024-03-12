package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBankAccount(context *gin.Context) {
	// db := database.GetDB()

	// userData := ctx.MustGet("userData").(jwt5.MapClaims)
	// contentType := helpers.GetContentType(ctx)

	// Book := models.Book{}
	// userID := uint(userData["id"].(float64))

	// if contentType == appJSON {
	// 	ctx.ShouldBindJSON(&Book)
	// } else {
	// 	ctx.ShouldBind(&Book)
	// }

	// Book.UserID = userID
	// newUUID := uuid.New()
	// Book.UUID = newUUID.String()

	// err := db.Debug().Create(&Book).Error
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{
	// 		"error":   "Bad request",
	// 		"message": err.Error(),
	// 	})
	// 	return
	// }

	context.JSON(http.StatusOK, gin.H{
		"data": "",
	})
}
