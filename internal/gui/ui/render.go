package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

const (
	preferenceCurrentPanel = "currentPanel"
	homePanel              = "home"
	aboutPanel             = "about"
	configurationPanel     = "configuration"
)

func Render() {
	a := app.NewWithID("com.jmillerv.daily")
	w := a.NewWindow("Analect - a cross-platform app for preserving quotes")

	w.SetMaster()

	w.Resize(fyne.Size{Width: 800, Height: 560})

	w.SetContent(createMainContainer())
	w.ShowAndRun()
}
