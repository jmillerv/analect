package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
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
	quoteTab := container.NewTabItem("Quote", createForm())
	archiveTab := container.NewTabItem("Archive", createArchive())
	aboutTab := container.NewTabItem("About", createAbout())

	tabContainer := container.NewAppTabs(quoteTab, archiveTab, aboutTab)

	return tabContainer
}
