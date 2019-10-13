package awss3

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Lists the items in the specified S3 Bucket
func ListS3Files(bucket string, prefix string) (*s3.ListObjectsOutput, error) {

	sess, err := AwsSession()
	if err != nil {
		return nil, err
	}

	// Create S3 service client
	svc := s3.New(sess)

	// Get the list of items
	// delimiter := "/"
	resp, err := svc.ListObjects(&s3.ListObjectsInput{
		Bucket: aws.String(bucket),
		Prefix: aws.String(prefix)},
	// Delimiter: aws.String("/")},
	)
	if err != nil {
		return nil, err
		//exitErrorf("Unable to list items in bucket %q, %v", bucket, err)
	}
	// logging.Info.Printf("Found %v items in %v/%v\n", len(resp.Contents), bucket, prefix)

	return resp, nil
}

func ListS3Prefixes(bucket string, prefix string) ([]string, error) {

	if prefix[len(prefix)-1:] != "/" {
		prefix = fmt.Sprintf("%v/", prefix)
	}

	sess, err := AwsSession()
	if err != nil {
		return nil, err
	}

	// Create S3 service client
	svc := s3.New(sess)

	resp, err := svc.ListObjects(&s3.ListObjectsInput{
		Bucket:    aws.String(bucket),
		Prefix:    aws.String(prefix),
		Delimiter: aws.String("/")},
	)
	if err != nil {
		return nil, err
	}
	// logging.Info.Printf("Found %v prefixes in %v/%v\n", len(resp.CommonPrefixes), bucket, prefix)
	var prefixes []string
	for _, prefix := range resp.CommonPrefixes {
		prefixes = append(prefixes, *prefix.Prefix)

	}
	return prefixes, nil
}
