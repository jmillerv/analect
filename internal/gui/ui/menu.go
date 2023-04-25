package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/jmillerv/analect/internal/gui/panels"
)

func createSideMenu(mainContent *fyne.Container) fyne.CanvasObject {
	addQuoteBtn := widget.NewButton("Add Quote", func() {
		mainContent.Objects = []fyne.CanvasObject{panels.QuoteForm()}
		mainContent.Refresh()
	})

	archiveBtn := widget.NewButton("Archive", func() {
		mainContent.Objects = []fyne.CanvasObject{panels.Archive()}
		mainContent.Refresh()
	})

	aboutBtn := widget.NewButton("About", func() {
		mainContent.Objects = []fyne.CanvasObject{panels.About()}
		mainContent.Refresh()
	})

	return container.NewVBox(addQuoteBtn, archiveBtn, aboutBtn)
}
