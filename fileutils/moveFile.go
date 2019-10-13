package fileutils

import (
	"fmt"
	"math/rand"
	"os"
	"path"
	"time"

	"github.com/pedrovog/gocommon_lib/logging"
)

func MoveFile(file string, destFolder string) error {
	err := CreateDir(destFolder)
	if err != nil {
		return err
	}
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	newFileName := fmt.Sprintf("%d-%v", seededRand.Int(), path.Base(file))
	destFile := path.Join(path.Dir(destFolder), newFileName)
	logging.Info.Printf("Moving file %v to %v: ", file, destFile)
	err = os.Rename(file, destFile)
	if err != nil {
		return err
	}
	return nil
}
