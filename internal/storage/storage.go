package storage

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

// saveDataToFile persists the data to the given filepath
func saveDataToFile(data []byte, path string) error {
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

// LoadData retrieves the data from the quotes.json file. This function assumes that a `saveDestination` has been set in the app preferences.
func LoadData(w fyne.Window) ([]byte, error) {
	// Retrieve file location from preferences
	prefs := fyne.CurrentApp().Preferences()
	path := prefs.String("saveDestination")

	if path == "" {
		err := fmt.Errorf("file location not set")
		dialog.ShowError(err, w)
		return nil, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		dialog.ShowError(fmt.Errorf("failed to read file: %v", err), w)
		return nil, err
	}
	return data, nil
}

func SaveData(w fyne.Window, data []byte) {
	// Retrieve the file location from preferences
	prefs := fyne.CurrentApp().Preferences()
	path := prefs.String("saveDestination")

	// if path is empty, throw a dialog error
	if path == "" {
		dialog.ShowError(fmt.Errorf("file location not set"), w)
		return
	}

	err := saveDataToFile(data, path)
	if err != nil {
		dialog.ShowError(err, w)
	} else {
		dialog.ShowInformation("Success", "Data saved successfully", w)
	}
}
