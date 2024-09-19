package awssdk

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var awsSession *session.Session

// AwsSessionInit initializes the AWS session using credentials and region from environment variables.
func AwsSessionInit() {
	// Load AWS configuration from environment variables
	region := os.Getenv("REGION")
	accessKey := os.Getenv("ACCESS_KEY_ID")
	secretKey := os.Getenv("SECRET_ACCESS_KEY")

	// Create a new AWS session
	sess, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region: aws.String(region),
			Credentials: credentials.NewStaticCredentials(
				accessKey,
				secretKey,
				"", // Session token can be provided here if needed
			),
		},
	})
	if err != nil {
		panic(fmt.Sprintf("Failed to initialize AWS session: %v", err))
	}

	awsSession = sess
	fmt.Println("AWS Session initialized")
}

// SaveFileS3 uploads a file to AWS S3 and returns the file URL.
func SaveFileS3(fileReader io.Reader, fileHeader *multipart.FileHeader) (string, error) {
	// Load bucket name from environment variables
	bucketName := os.Getenv("BUCKET_NAME")

	// Create a new uploader instance
	uploader := s3manager.NewUploader(awsSession)

	// Generate a unique file key using the current timestamp
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	fileKey := fmt.Sprintf("%d_%s", timestamp, fileHeader.Filename)

	// Upload the file to S3
	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileKey),
		Body:   fileReader,
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload file to S3: %w", err)
	}

	// Generate the URL of the uploaded file
	url := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucketName, fileKey)
	return url, nil
}
