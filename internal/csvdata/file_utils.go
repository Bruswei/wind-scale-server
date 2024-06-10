package csvdata

import (
	"os"
)

func FileExists(filePath string) error {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer file.Close()

		if _, err := file.WriteString("Location;Time;WindSpeed\n"); err != nil {
			return err
		}
	}
	return err
}
