package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/jmillerv/analect/internal/models"
)

var quotes []models.Quote

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
