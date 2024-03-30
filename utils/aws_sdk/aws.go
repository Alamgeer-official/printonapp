package awssdk

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var awsSession *session.Session

func AwsSessionInit() {
	var region string = os.Getenv("REGION")
	var accessKey string = os.Getenv("ACCESS_KEY_ID")
	var secretKey string = os.Getenv("SECRET_ACCESS_KEY")
	fmt.Printf("region: %v\n", region)

	Session, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region: aws.String(region),
			Credentials: credentials.NewStaticCredentials(
				accessKey,
				secretKey,
				"",
			),
		},
	})

	if err != nil {
		panic(err)
	}

	awsSession = Session

	fmt.Println("AWS Session initialized")

}

func SaveFileS3(fileReader io.Reader, fileHeader *multipart.FileHeader) (string, error) {
	
	var bucketName string = os.Getenv("BUCKET_NAME")

	//get session into upload file
	uploader := s3manager.NewUploader(awsSession)
	// Upload the file to S3 using the fileReader
	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileHeader.Filename),
		Body:   fileReader,
	})
	if err != nil {
		return "", err
	}

	// Get the URL of the uploaded file
	url := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucketName, fileHeader.Filename)

	return url, nil
}
