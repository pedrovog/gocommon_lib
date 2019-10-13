package awss3

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/pedrovog/gocommon_lib/fileutils"
	"github.com/pedrovog/gocommon_lib/logging"
)

func DownloadS3File(bucket string, itemKey string) (string, error) {

	err := fileutils.CreateDir(itemKey)
	if err != nil {
		return "", err
	}

	logging.Info.Println("Downloading File: ", itemKey)
	file, err := os.Create(itemKey)
	if err != nil {
		return "", err
	}
	defer file.Close()
	sess, err := AwsSession()
	if err != nil {
		return "", err
	}
	downloader := s3manager.NewDownloader(sess)

	numBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(itemKey),
		})
	if err != nil {
		return "", err
	}

	logging.Info.Println("Downloaded File: ", file.Name(), numBytes, "bytes")
	return file.Name(), nil
}

func DownloadS3FileV2(bucket string, itemKey string, destination string) (string, error) {

	err := fileutils.CreateDir(destination)
	if err != nil {
		return "", err
	}

	logging.Info.Println("Downloading File: ", itemKey)
	file, err := os.Create(destination)
	if err != nil {
		return "", err
	}
	defer file.Close()
	sess, err := AwsSession()
	if err != nil {
		return "", err
	}
	downloader := s3manager.NewDownloader(sess)

	numBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(itemKey),
		})
	if err != nil {
		return "", err
	}

	logging.Info.Println("Downloaded File: ", file.Name(), numBytes, "bytes")
	return file.Name(), nil
}
