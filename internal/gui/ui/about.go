package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func createAbout() fyne.CanvasObject {
	// Create the widgets to display information
	title := widget.NewLabel("Analect")
	version := widget.NewLabel("Version: 1.0")
	author := widget.NewLabel("Jeremiah Miller")
	description := widget.NewLabelWithStyle("This application allows you to add and manage quotes in a JSON format. You can add new quotes and view the archive of saved quotes.",
		fyne.TextAlignLeading, fyne.TextStyle{})
	// Add any icons or images you'd like to display
	appIcon := widget.NewIcon(theme.DocumentIcon()) // Replace theme.DocumentIcon() with the path to your custom icon

	// Customize the appearance of the widgets
	title.TextStyle = fyne.TextStyle{Bold: true, Italic: true}
	version.TextStyle = fyne.TextStyle{Italic: true}
	author.TextStyle = fyne.TextStyle{Italic: true}

	// Arrange the widgets within a container
	aboutContent := container.NewVBox(
		title,
		version,
		author,
		appIcon,
		description,
	)

	// Add padding around the content
	aboutContainer := container.NewPadded(aboutContent)

	return aboutContainer
}
