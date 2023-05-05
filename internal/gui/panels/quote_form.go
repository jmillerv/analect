package panels

import (
	"encoding/json"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/jmillerv/analect/internal/models"
	"github.com/jmillerv/analect/internal/storage"
)

func QuoteForm(w fyne.Window) fyne.CanvasObject {
	authorEntry := widget.NewEntry()
	quoteEntry := widget.NewEntry()
	citationEntry := widget.NewEntry()
	linkEntry := widget.NewEntry()

	saveButton := widget.NewButton("Save", func() {
		saveJSONData(w, authorEntry.Text, quoteEntry.Text, citationEntry.Text, linkEntry.Text)
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

func saveJSONData(w fyne.Window, author, quote, citation, link string) {
	// load existing data
	quoteList, err := loadQuotes(w)
	if err != nil {
		dialog.ShowError(err, w)
	}

	// Create a new QuoteObject and append it to the JSON array
	quoteObj := models.Quote{Author: author, Quote: quote, Citation: citation, Link: link}
	quoteList.AddQuote(&quoteObj)
	data, err := json.MarshalIndent(quoteList, "", "  ") // marshall the json in a more readable format
	if err != nil {
		log.Println("Error marshaling JSON:", err)
		return
	}

	// Save the JSON array to the file
	storage.SaveData(w, data)
}

func loadQuotes(w fyne.Window) (*models.QuoteList, error) {
	data, err := storage.LoadData(w)
	if err != nil {
		return nil, err
	}

	var quotes models.QuoteList
	err = json.Unmarshal(data, &quotes)
	if err != nil {
		return nil, err
	}

	return &quotes, nil
}
