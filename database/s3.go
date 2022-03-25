//for storage
package database

import (
	"io"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func ConnectS3(keyId string, appKey string, Endpoint string, region string) *session.Session {
	s3Config := aws.Config{
		Credentials:      credentials.NewStaticCredentials(keyId, appKey, ""),
		Endpoint:         aws.String(Endpoint),
		Region:           aws.String(region),
		S3ForcePathStyle: aws.Bool(true),
	}
	session, err := session.NewSession(&s3Config)
	if err != nil {
		log.Panicln("S3 line:22")
		log.Fatal(err)
	}

	return session
}

func CreateS3Bucket(session *session.Session, bucketName string) error {
	cpram := &s3.CreateBucketInput{
		Bucket: &bucketName,
	}
	client := s3.New(session)
	_, err := client.CreateBucket(cpram)
	return err
}

func CreateUploader(session *session.Session) *s3manager.Uploader {
	return s3manager.NewUploader(session)
}

func InsertObjectWithS3(uploader *s3manager.Uploader, bucketName string, key string, content io.Reader) error {
	output, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
		Body:   content,
	})
	log.Println(output)
	return err
}
