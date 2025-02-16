package handlers

import (
	"github.com/gin-gonic/gin"
	"go-file-storage/config"
	"net/http"
	"github.com/aws/aws-sdk-go/service/s3"
)

func GetFile(c *gin.Context) {
	filename := c.Param("id")

	req, _ := config.S3Client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(config.BucketName),
		Key:    aws.String(filename),
	})

	url, err := req.Presign(15 * 60)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate file URL"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"download_url": url})
}
