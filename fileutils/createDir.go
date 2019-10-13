package fileutils

import (
	"os"
	"path"

	"github.com/pedrovog/gocommon_lib/logging"
)

func CreateDir(directory string) error {
	if _, err := os.Stat(path.Dir(directory)); os.IsNotExist(err) {
		logging.Info.Println("Creating Directory: ", path.Dir(directory))
		err := os.MkdirAll(path.Dir(directory), 0777)
		if err != nil {
			logging.Info.Printf("ERROR ao CRIAR diretorio %v -- %v\n", path.Dir(directory), err.Error())
			return err
		}
	}
	return nil
}
