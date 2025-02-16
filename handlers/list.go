package handlers

import (
	"github.com/gin-gonic/gin"
	"go-gin-app/config"
	"net/http"
	"github.com/aws/aws-sdk-go/service/s3"
)

func ListFiles(c *gin.Context) {
	resp, err := config.S3Client.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: &config.BucketName,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list files"})
		return
	}

	files := []string{}
	for _, item := range resp.Contents {
		files = append(files, *item.Key)
	}

	c.JSON(http.StatusOK, gin.H{"files": files})
}
