package awss3

import (
	"fmt"
)

func FileExists(bucketName string, filePath string) (bool, error) {
	f, err := ListS3Files(bucketName, filePath)
	if err != nil {
		fmt.Println(err)
	}
	if len(f.Contents) == 1 {
		return true, nil
	}
	return false, nil
}
