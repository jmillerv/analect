package storage

import "github.com/jmillerv/analect/models"

type Saver interface {
	Save(data models.QuoteList) error
}

type Loader interface {
	Load() (data models.QuoteList, err error)
}