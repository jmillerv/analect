package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/jmillerv/analect/models"
)

var quotes []models.Quote

func main() {
	myApp := app.New()
	myApp.Settings().SetTheme(theme.DarkTheme())

	w := myApp.NewWindow("Analect")
	w.SetFixedSize(false) // Allows resizing
	w.Resize(fyne.NewSize(600, 400))
	w.SetContent(createMainContainer())
	w.ShowAndRun()
}

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

func createArchive() fyne.CanvasObject {
	searchBar := widget.NewEntry()
	searchBar.SetPlaceHolder("Search quotes")

	quoteList := widget.NewList(
		func() int {
			// return the number of quotes
			return len(quotes)
		},
		func() fyne.CanvasObject {
			// return a quote item template
			return widget.NewLabel("")
		},
		func(i widget.ListItemID, c fyne.CanvasObject) {
			// update the content of a quote item
			c.(*widget.Label).SetText(quotes[i].Quote)
		},
	)

	searchBar.OnChanged = func(s string) {
		// filter the quotes based on the search term
		// filteredQuotes := searchQuotes(s)
		// quoteList.Length(len(filteredQuotes.Quotes))
		quoteList.Refresh()
	}

	archiveContainer := container.NewVBox(searchBar, quoteList)

	return archiveContainer
}

func searchQuotes(searchTerm string) *models.QuoteList {
	// var filteredQuotes []models.QuoteList
	//
	// for _, quote := range quotes {
	// 	if strings.Contains(strings.ToLower(quote.Quote), strings.ToLower(searchTerm)) {
	// 		filteredQuotes = append(filteredQuotes, quote)
	// 	}
	// }
	//
	// return filteredQuotes
	return nil
}

func createForm() fyne.CanvasObject {
	authorEntry := widget.NewEntry()
	quoteEntry := widget.NewEntry()
	citationEntry := widget.NewEntry()
	linkEntry := widget.NewEntry()

	saveButton := widget.NewButton("Save", func() {
		saveJSONData(authorEntry.Text, quoteEntry.Text, citationEntry.Text, linkEntry.Text)
	})

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Author", Widget: authorEntry},
			{Text: "Quote", Widget: quoteEntry},
			{Text: "Citation", Widget: citationEntry},
			{Text: "Link", Widget: linkEntry},
		},
		OnSubmit:   saveButton.OnTapped,
		SubmitText: "Save",
	}

	return container.NewVBox(form)
}

func saveJSONData(author, quote, citation, link string) {
	// Load the data from dropbox

	// Create a new QuoteObject and append it to the JSON array
	// quoteObj := models.Quote{Author: author, Quote: quote, Citation: citation, Link: link}
	// data, err := json.Marshal(quoteObj)
	// if err != nil {
	// 	log.Println("Error marshaling JSON:", err)
	// 	return
	// }

	// Save the JSON array to Dropbox
	// saveToDropbox(data)
}

type FixedSideMenuLayout struct {
	MenuWidth int
}

func (l *FixedSideMenuLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	if len(objects) != 2 {
		return
	}

	menuSize := fyne.NewSize(float32(l.MenuWidth), size.Height)
	contentSize := fyne.NewSize(float32(size.Width)-float32(l.MenuWidth), size.Height)

	objects[0].Resize(menuSize)
	objects[0].Move(fyne.NewPos(0, 0))

	objects[1].Resize(contentSize)
	objects[1].Move(fyne.NewPos(float32(l.MenuWidth), 0))
}

func (l *FixedSideMenuLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(0, 0)
}

func createTabContainer() *container.AppTabs {
	quoteTab := container.NewTabItem("Quote", createForm())
	archiveTab := container.NewTabItem("Archive", createArchive())
	aboutTab := container.NewTabItem("About", createAbout())

	tabContainer := container.NewAppTabs(quoteTab, archiveTab, aboutTab)

	return tabContainer
}
