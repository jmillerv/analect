package storage

import "github.com/jmillerv/analect/internal/models"

type Local struct {
	Path string
}

func (*Local) Save(data *models.QuoteList) error {
	return nil
}

func (*Local) Load() (data *models.QuoteList, err error) {
	return &models.QuoteList{}, nil
}
