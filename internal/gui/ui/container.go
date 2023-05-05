package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/jmillerv/analect/internal/gui/panels"
)

func createMainContainer(window fyne.Window) fyne.CanvasObject {
	mainContent := container.NewMax()
	sideMenu := createSideMenu(mainContent, window)

	if fyne.CurrentDevice().IsMobile() {
		tabContainer := createTabContainer(window)
		return tabContainer
	} else {
		return container.NewBorder(nil, nil, sideMenu, nil, mainContent)
	}
}

func createTabContainer(w fyne.Window) *container.AppTabs {
	quoteTab := container.NewTabItem("Quote", panels.QuoteForm(w))
	archiveTab := container.NewTabItem("Archive", panels.Archive(w))
	cloudTab := container.NewTabItem("Cloud Settings", nil)
	aboutTab := container.NewTabItem("About", panels.About())

	tabContainer := container.NewAppTabs(quoteTab, archiveTab, cloudTab, aboutTab)

	return tabContainer
}
