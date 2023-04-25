package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func createSideMenu(mainContent *fyne.Container) fyne.CanvasObject {
	addQuoteBtn := widget.NewButton("Add Quote", func() {
		mainContent.Objects = []fyne.CanvasObject{createForm()}
		mainContent.Refresh()
	})

	archiveBtn := widget.NewButton("Archive", func() {
		mainContent.Objects = []fyne.CanvasObject{createArchive()}
		mainContent.Refresh()
	})

	aboutBtn := widget.NewButton("About", func() {
		mainContent.Objects = []fyne.CanvasObject{createAbout()}
		mainContent.Refresh()
	})

	return container.NewVBox(addQuoteBtn, archiveBtn, aboutBtn)
}
