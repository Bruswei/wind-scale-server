package csvdata

import (
	"os"
	"path/filepath"
)

func FileExists(filePath string) error {
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer file.Close()

		if _, err := file.WriteString("Location,Time,WindSpeed\n"); err != nil {
			return err
		}
		return nil
	}
	return err
}
