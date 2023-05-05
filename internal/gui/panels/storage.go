package panels

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

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
