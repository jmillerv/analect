package panels

import (
	"net/url"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/jmillerv/analect/internal/models"
	"github.com/sirupsen/logrus"
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

			authorLabel := widget.NewLabel("")
			authorLabel.Wrapping = fyne.TextWrapWord

			// quoteLabel := widget.NewLabel("")
			// quoteLabel.Wrapping = fyne.TextTruncate

			quoteButton := widget.NewButton("", func() {})
			quoteButton.ExtendBaseWidget(quoteButton)

			linkLabel := widget.NewHyperlink("", nil)
			linkLabel.Wrapping = fyne.TextWrapWord

			return container.New(layout.NewGridLayoutWithColumns(3), authorLabel, quoteButton, linkLabel)
		},
		func(i widget.ListItemID, c fyne.CanvasObject) {
			// update the content of a quote item
			grid := c.(*fyne.Container)
			grid.Objects[0].(*widget.Label).SetText(filteredQuotes.Quotes[i].Author)
			quoteButton := grid.Objects[1].(*widget.Button)
			quoteButton.SetText(getFirst50Chars(filteredQuotes.Quotes[i].Quote))
			quoteButton.SetIcon(theme.MenuExpandIcon())
			quoteButton.OnTapped = func() {
				dialog := createQuotePopOut(w, filteredQuotes.Quotes[i])
				dialog.Show()
			}
			grid.Objects[2].(*widget.Hyperlink).SetText(filteredQuotes.Quotes[i].Link)
			linkLabel := grid.Objects[2].(*widget.Hyperlink)
			linkLabel.SetText(filteredQuotes.Quotes[i].Citation)
			err := linkLabel.SetURLFromString(filteredQuotes.Quotes[i].Link)
			if err != nil {
				logrus.WithError(err).Error("unable to set link")
			}
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

func createQuotePopOut(w fyne.Window, quote models.Quote) fyne.CanvasObject {
	var modal *widget.PopUp

	quoteHyperlink, err := url.Parse(quote.Link)
	if err != nil {
		logrus.WithError(err).Error("unable to parse quote link")
	}
	authorLabel := widget.NewLabel(quote.Author)
	quoteLabel := widget.NewLabel(quote.Quote)
	quoteLabel.Wrapping = fyne.TextWrapWord
	linkLabel := widget.NewHyperlink(quote.Citation, quoteHyperlink)
	hideModal := widget.NewButton("Close", func() {
		if modal != nil {
			modal.Hide()
		}
	})
	// Set a minimum size for the quote label and wrap it with a ScrollContainer
	scroll := container.NewScroll(quoteLabel)
	scroll.SetMinSize(fyne.NewSize(400, 300))

	content := container.NewVBox(
		authorLabel,
		scroll,
		linkLabel,
		hideModal,
	)

	modal = widget.NewModalPopUp(content, w.Canvas())
	modal.Show()
	return modal
}

func getFirst50Chars(s string) string {
	if len(s) > 25 {
		return s[:25]
	}
	return s
}
