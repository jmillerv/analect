package models

type Quote struct {
	Author   string   `json:"author"`
	Quote    string   `json:"quote"`
	Citation string   `json:"citation"`
	Link     string   `json:"link"`
	Tags     []string `json:"tags"`
}

type QuoteList struct {
	Quotes []Quote `json:"quotes"`
}

func (ql *QuoteList) AddQuote(quote *Quote) {
	ql.Quotes = append(ql.Quotes, *quote)
}
