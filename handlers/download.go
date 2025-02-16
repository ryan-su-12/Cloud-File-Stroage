package handlers

import (
	"github.com/gin-gonic/gin"
	"go-gin-app/config"
	"net/http"
	"github.com/aws/aws-sdk-go/service/s3"

)

func GetFile(c *gin.Context) {  // ✅ Only this function should be here
	filename := c.Param("id")

	req, _ := config.S3Client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: &config.BucketName,
		Key:    &filename,
	})

	url, err := req.Presign(15 * 60)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate file URL"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"download_url": url})
}

