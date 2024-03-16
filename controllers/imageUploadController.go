package controllers

import (
	// "fmt"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/olahol/go-imageupload"
)

func CreateUploadImage(ctx *gin.Context) {
	contentType := ctx.GetHeader("Content-Type")
	// fmt.Println(contentType)
	if !strings.HasPrefix(contentType, "multipart/form-data") {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Content-Type must be multipart/form-data"})
		return
	}

	// Process uploaded image
	img, err := imageupload.Process(ctx.Request, "file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to process uploaded image"})
		return
	}

	// Check file size
	if img.Size > 2*1024*1024 || img.Size < 10*1024 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Uploaded file size must be between 10KB and 2MB"})
		return
	}

	// Check file extension
	if !strings.HasSuffix(strings.ToLower(img.Filename), ".jpg") && !strings.HasSuffix(strings.ToLower(img.Filename), ".jpeg") {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Uploaded file must be in *.jpg or *.jpeg format"})
		return
	}

	// Generate filename
	filename := uuid.New().String() + ".jpeg"

	dir, err := os.Getwd()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "GETWD error"})
		return
	}
	// fmt.Println(dir)

	// Save image
	err = img.Save(dir + "/uploads/" + filename)
	// fmt.Println(err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save uploaded image"})
		return
	}

	// Upload image to AWS S3
	err = uploadToS3(dir+"/uploads/"+filename, filename)
	// fmt.Println(err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload image to AWS S3"})
		return
	}

	// Construct image URL
	imageURL := "https://s3.amazonaws.com/" + os.Getenv("S3_BUCKET_NAME") + "/" + filename

	// Respond with image URL
	ctx.JSON(http.StatusOK, gin.H{"imageUrl": imageURL})
}

// Uploads file to AWS S3
func uploadToS3(filePath, filename string) error {
	// Create AWS session
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("ap-southeast-1"),
		Credentials: credentials.NewStaticCredentials(os.Getenv("S3_ID"), os.Getenv("S3_SECRET_KEY"), ""),
	})
	if err != nil {
		return err
	}

	// Create S3 service client
	svc := s3.New(sess)

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Upload file to S3
	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("S3_BUCKET_NAME")),
		Key:    aws.String(filename),
		ACL:    aws.String("public-read"),
		Body:   file,
	})
	if err != nil {
		return err
	}

	return nil
}
