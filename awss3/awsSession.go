package awss3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

var (
	// AwsRegion Region that the session will be acquired
	AwsRegion = "us-east-2"
)

// AwsSession Acquire an aws session on specified AwsRegion
func AwsSession() (*session.Session, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(AwsRegion)},
	)
	if err != nil {
		return nil, err
	}
	return sess, nil
}
