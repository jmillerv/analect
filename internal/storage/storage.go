package storage

import "os"

// SaveDataToFile persists the data to the given filepath
func SaveDataToFile(data []byte, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}
