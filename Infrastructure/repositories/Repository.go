package repositories

import (
	"github.com/apmath-web/currency/Domain"
)

type Repository struct {
	currentRates Domain.ChangeTable
}

func (repository Repository) SetAllTable(table Domain.ChangeTable) error {
	repository.currentRates = table
	return nil
}

func (repository Repository) GetAllTable() Domain.ChangeTable {
	return repository.currentRates
}

func (repository Repository) ClearAllTable() error {
	repository.currentRates = nil
	return nil
}
