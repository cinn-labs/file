package file

import (
	"os"

	"github.com/cinn-labs/lg"
)

func FromString(filePath, s string) error {
	f, err := os.Create(filePath)
	if lg.CError(err) {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(s)
	if lg.CError(err) {
		return err
	}

	return err
}

func FromStringWithoutOverriding(filePath, s string) error {
	var err error
	if _, err = os.Stat(filePath); os.IsNotExist(err) {
		return FromString(filePath, s)
	}
	return err
}
