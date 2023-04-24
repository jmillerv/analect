package storage

import "github.com/jmillerv/analect/models"

func Save(data *models.QuoteList) error {
	return nil
}

func Load() (*models.QuoteList, error) {
	return &models.QuoteList{}, nil
}
