package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"github.com/jmillerv/analect/internal/gui/panels"
)

func createMainContainer() fyne.CanvasObject {
	mainContent := container.New(layout.NewVBoxLayout())
	sideMenu := createSideMenu(mainContent)

	if fyne.CurrentDevice().IsMobile() {
		tabContainer := createTabContainer()
		return tabContainer
	} else {
		return container.New(&FixedSideMenuLayout{MenuWidth: 200},
			sideMenu,
			mainContent,
		)
	}
}

func createTabContainer() *container.AppTabs {
	quoteTab := container.NewTabItem("Quote", panels.QuoteForm())
	archiveTab := container.NewTabItem("Archive", panels.Archive())
	aboutTab := container.NewTabItem("About", panels.About())

	tabContainer := container.NewAppTabs(quoteTab, archiveTab, aboutTab)

	return tabContainer
}
