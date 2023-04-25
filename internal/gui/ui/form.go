package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

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
