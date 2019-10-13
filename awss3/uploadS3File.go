package awss3

import (
	"bytes"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/pedrovog/gocommon_lib/logging"
)

// Upload file to s3 bucket
func UploadS3File(bucket string, file string) error {

	// Open the file for use
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	// Get file size and read the file content into a buffer
	fileInfo, _ := f.Stat()
	size := fileInfo.Size()
	buffer := make([]byte, size)
	f.Read(buffer)

	sess, err := AwsSession()
	if err != nil {
		return err
	}

	logging.Info.Printf("Uploading file %v to bucket %v\n", file, bucket)
	// Config settings: this is where you choose the bucket, filename, content-type etc.
	// of the file you're uploading.
	_, err = s3.New(sess).PutObject(&s3.PutObjectInput{
		Bucket:        aws.String(bucket),
		Key:           aws.String(file),
		Body:          bytes.NewReader(buffer),
		ContentLength: aws.Int64(size),
		ContentType:   aws.String(http.DetectContentType(buffer)),
	})
	if err != nil {
		return err
	}
	logging.Info.Printf("File %v uploaded to bucket %v with success\n", file, bucket)
	return nil
}

// Upload file to s3 bucket
func UploadS3FileV2(bucket string, file string, destination string) error {

	// Open the file for use
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	// Get file size and read the file content into a buffer
	fileInfo, _ := f.Stat()
	size := fileInfo.Size()
	buffer := make([]byte, size)
	f.Read(buffer)

	sess, err := AwsSession()
	if err != nil {
		return err
	}

	logging.Info.Printf("Uploading file %v to bucket %v\n", file, bucket)
	// Config settings: this is where you choose the bucket, filename, content-type etc.
	// of the file you're uploading.
	_, err = s3.New(sess).PutObject(&s3.PutObjectInput{
		Bucket:        aws.String(bucket),
		Key:           aws.String(destination),
		Body:          bytes.NewReader(buffer),
		ContentLength: aws.Int64(size),
		ContentType:   aws.String(http.DetectContentType(buffer)),
	})
	if err != nil {
		return err
	}
	logging.Info.Printf("File %v uploaded to bucket %v with success\n", file, bucket)
	return nil
}
