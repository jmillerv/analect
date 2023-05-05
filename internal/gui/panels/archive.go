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

			// TODO investigate if a fyne table may be better for this UI
			authorLabel := widget.NewLabel("")
			authorLabel.Wrapping = fyne.TextWrapWord

			quoteLabel := widget.NewLabel("")
			quoteLabel.Wrapping = fyne.TextTruncate

			citationLabel := widget.NewLabel("")
			citationLabel.Wrapping = fyne.TextWrapOff

			linkLabel := widget.NewLabel("")
			linkLabel.Wrapping = fyne.TextWrapOff

			return container.NewGridWithColumns(4, authorLabel, quoteLabel, citationLabel, linkLabel)
		},
		func(i widget.ListItemID, c fyne.CanvasObject) {
			// update the content of a quote item
			grid := c.(*fyne.Container)
			grid.Objects[0].(*widget.Label).SetText(filteredQuotes.Quotes[i].Author)
			grid.Objects[1].(*widget.Label).SetText(filteredQuotes.Quotes[i].Quote)
			grid.Objects[2].(*widget.Label).SetText(filteredQuotes.Quotes[i].Citation)
			grid.Objects[3].(*widget.Label).SetText(filteredQuotes.Quotes[i].Link)
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
	quoteList.Resize(fyne.NewSize(200, 400))

	archiveContainer := container.NewBorder(searchBar, nil, nil, nil, quoteList)

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
