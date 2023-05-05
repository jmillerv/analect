package panels

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"github.com/jmillerv/analect/internal/storage"
)

func saveData(w fyne.Window, data []byte) {
	// Retrieve the file location from preferences
	prefs := fyne.CurrentApp().Preferences()
	path := prefs.String("saveDestination")

	// if path is empty, throw a dialog error
	if path == "" {
		dialog.ShowError(fmt.Errorf("file location not set"), w)
		return
	}

	err := storage.SaveDataToFile(data, path)
	if err != nil {
		dialog.ShowError(err, w)
	} else {
		dialog.ShowInformation("Success", "Data saved successfully", w)
	}
}

func ShowFileLocationDialog(w fyne.Window) {
	locationDialog := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
		if err == nil && writer != nil {
			defer writer.Close()

			path := writer.URI().Path()
			// Save the selected location to preferences
			prefs := fyne.CurrentApp().Preferences()
			prefs.SetString("saveDestination", path)
		}
	}, w)

	locationDialog.SetFileName("quotes.json") // Set the suggested file name to "quotes.json"
	locationDialog.Show()
}
