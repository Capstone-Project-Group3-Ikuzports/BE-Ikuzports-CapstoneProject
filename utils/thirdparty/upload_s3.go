package thirdparty

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

// CREATE RANDOM STRING

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func autoGenerate(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func String(length int) string {
	return autoGenerate(length, charset)
}

// UPLOAD TO AWS S3

func UploadProfile(c echo.Context, image string) (string, error) {

	file, fileheader, err := c.Request().FormFile(image)
	if err != nil {
		log.Print(err)
		return "", err
	}

	randomStr := String(20)

	godotenv.Load(".env")

	fmt.Println("aws region", os.Getenv("AWS_REGION"))
	fmt.Println("access key", os.Getenv("ACCESS_KEY_IAM"))
	fmt.Println("SECRET_KEY_IAM", os.Getenv("SECRET_KEY_IAM"))
	fmt.Println("AWS_BUCKET_NAME", os.Getenv("AWS_BUCKET_NAME"))

	s3Config := &aws.Config{
		Region:      aws.String(os.Getenv("AWS_REGION")),
		Credentials: credentials.NewStaticCredentials(os.Getenv("ACCESS_KEY_IAM"), os.Getenv("SECRET_KEY_IAM"), ""),
	}
	s3Session := session.New(s3Config)

	uploader := s3manager.NewUploader(s3Session)

	input := &s3manager.UploadInput{
		Bucket:      aws.String(os.Getenv("AWS_BUCKET_NAME")),                           // bucket's name
		Key:         aws.String("testetstets/" + randomStr + "-" + fileheader.Filename), // files destination location
		Body:        file,                                                               // content of the file
		ContentType: aws.String("image/jpg"),                                            // content type
	}
	res, err := uploader.UploadWithContext(context.Background(), input)

	// RETURN URL LOCATION IN AWS
	return res.Location, err
}

func Upload(c echo.Context) (string, error) {

	file, fileheader, err := c.Request().FormFile("file")
	if err != nil {
		log.Print(err)
		return "", err
	}

	randomStr := String(20)

	godotenv.Load(".env")

	s3Config := &aws.Config{
		Region:      aws.String(os.Getenv("AWS_REGION")),
		Credentials: credentials.NewStaticCredentials(os.Getenv("ACCESS_KEY_IAM"), os.Getenv("SECRET_KEY_IAM"), ""),
	}
	s3Session := session.New(s3Config)

	uploader := s3manager.NewUploader(s3Session)

	input := &s3manager.UploadInput{
		Bucket:      aws.String(os.Getenv("AWS_BUCKET_NAME")),                        // bucket's name
		Key:         aws.String("property/" + randomStr + "-" + fileheader.Filename), // files destination location
		Body:        file,                                                            // content of the file
		ContentType: aws.String("image/jpg"),                                         // content type
	}
	res, err := uploader.UploadWithContext(context.Background(), input)

	// RETURN URL LOCATION IN AWS
	return res.Location, err
}
