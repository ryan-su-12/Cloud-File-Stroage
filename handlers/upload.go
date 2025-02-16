package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-app/config"
	"net/http"
	"strings"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws"  // ✅ Add this line
)

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file"})
		return
	}

	filename := strings.ReplaceAll(file.Filename, " ", "_")
	f, err := file.Open()
	defer f.Close()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
		return
	}

	_, err = config.S3Client.PutObject(&s3.PutObjectInput{
		Bucket: &config.BucketName,
		Key:    &filename,
		Body:   f,
		ACL:    aws.String("private"),
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file"})
		return
	}

	fileURL := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", config.BucketName, filename)

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully", "file_url": fileURL})
}
