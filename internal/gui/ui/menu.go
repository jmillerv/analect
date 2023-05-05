package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/jmillerv/analect/internal/gui/panels"
)

func createSideMenu(mainContent *fyne.Container, w fyne.Window) fyne.CanvasObject {
	addQuoteBtn := widget.NewButton("Add Quote", func() {
		mainContent.Objects = []fyne.CanvasObject{panels.QuoteForm(w)}
		mainContent.Refresh()
	})

	archiveBtn := widget.NewButton("Archive", func() {
		mainContent.Objects = []fyne.CanvasObject{panels.Archive()}
		mainContent.Refresh()
	})

	storageBtn := widget.NewButton("Storage", func() {
		panels.ShowFileLocationDialog(w)
	})

	aboutBtn := widget.NewButton("About", func() {
		mainContent.Objects = []fyne.CanvasObject{panels.About()}
		mainContent.Refresh()
	})

	return container.NewVBox(addQuoteBtn, archiveBtn, storageBtn, aboutBtn)
}
