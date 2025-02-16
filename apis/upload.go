package handlers

import (
	"github.com/gin-gonic/gin"
	"go-file-storage/config"
	"net/http"
	"strings"
	"github.com/aws/aws-sdk-go/service/s3"
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
		Bucket: aws.String(config.BucketName),
		Key:    aws.String(filename),
		Body:   f,
		ACL:    aws.String("private"),
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully", "file_url": "https://s3.amazonaws.com/" + config.BucketName + "/" + filename})
}
