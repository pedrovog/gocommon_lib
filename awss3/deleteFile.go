package awss3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

// DeleteFile deletes a S3 object
func DeleteFile(bucketName string, key string) (bool, error) {

	session, err := AwsSession()
	if err != nil {
		return false, err
	}

	svc := s3.New(session)

	input := &s3.DeleteObjectInput{Bucket: aws.String(bucketName), Key: aws.String(key)}

	_, err = svc.DeleteObject(input)
	if err != nil {
		return false, err
	}

	err = svc.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	})
	if err != nil {
		return false, err
	}

	return true, nil
}
