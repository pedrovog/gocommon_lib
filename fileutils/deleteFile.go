package fileutils

import (
	"os"

	"github.com/pedrovog/gocommon_lib/logging"
)

func DeleteFile(file string) error {
	logging.Info.Printf("Deleting file: %v ", file)
	err := os.Remove(file)
	if err != nil {
		return nil
	}
	logging.Info.Printf("File %v deleted", file)

	return nil
}
