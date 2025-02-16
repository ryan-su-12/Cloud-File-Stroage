package config

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/joho/godotenv"
)

var S3Client *s3.S3
var BucketName string

func InitAWS() {
	// Load .env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get bucket name
	BucketName = os.Getenv("AWS_BUCKET_NAME")

	// Create AWS session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	})
	if err != nil {
		log.Fatal("Failed to initialize AWS session:", err)
	}

	S3Client = s3.New(sess)

	// Test S3 connection by listing available buckets
	result, err := S3Client.ListBuckets(nil)
	if err != nil {
		log.Fatal("Failed to list S3 buckets:", err)
	}

	// Print bucket names
	fmt.Println("✅ Successfully connected to AWS S3! Available Buckets:")
	for _, b := range result.Buckets {
		fmt.Println(" -", *b.Name)
	}
}
