package handlers

import (
	"github.com/gin-gonic/gin"
	"go-file-storage/config"
	"net/http"
	"github.com/aws/aws-sdk-go/service/s3"
)

func DeleteFile(c *gin.Context) {
	filename := c.Param("id")

	_, err := config.S3Client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(config.BucketName),
		Key:    aws.String(filename),
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File deleted successfully"})
}
