package panels

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/jmillerv/analect/internal/models"
)

func Archive(w fyne.Window) fyne.CanvasObject {
	// Load the quotes from the quotes.json file
	quotes, err := loadQuotes(w)
	if err != nil {
		return widget.NewLabel("Error loading quotes")
	}

	searchBar := widget.NewEntry()
	searchBar.SetPlaceHolder("Search quotes")

	filteredQuotes := quotes

	quoteList := widget.NewList(
		func() int {
			// return the number of quotes
			return len(filteredQuotes.Quotes)
		},
		func() fyne.CanvasObject {
			// return a quote item template
			label := widget.NewLabel("")
			label.Wrapping = fyne.TextWrapWord
			return label
		},
		func(i widget.ListItemID, c fyne.CanvasObject) {
			// update the content of a quote item
			c.(*widget.Label).SetText(filteredQuotes.Quotes[i].Quote)
		},
	)
	updateList := func() {
		filteredQuotes = searchQuotes(quotes.Quotes, searchBar.Text)
		quoteList.Refresh()
	}

	searchBar.OnChanged = func(s string) {
		// filter the quotes based on the search term
		updateList()
	}

	listContainer := container.NewScroll(quoteList)
	archiveContainer := container.NewVBox(
		searchBar,
		listContainer,
	)

	return archiveContainer
}

func searchQuotes(quotes []models.Quote, searchTerm string) *models.QuoteList {
	if searchTerm == "" {
		return &models.QuoteList{
			Quotes: quotes,
		}
	}

	var filteredQuotes []models.Quote

	for _, quote := range quotes {
		if strings.Contains(strings.ToLower(quote.Quote), strings.ToLower(searchTerm)) {
			filteredQuotes = append(filteredQuotes, quote)
		}
	}

	return &models.QuoteList{
		Quotes: filteredQuotes,
	}
}
