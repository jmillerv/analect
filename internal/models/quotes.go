package models

type Quote struct {
	Author   string `json:"author"`
	Quote    string `json:"quote"`
	Citation string `json:"citation"`
	Link     string `json:"link"`
}

type QuoteList struct {
	Quotes []Quote `json:"quotes"`
}
