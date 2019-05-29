package repositories

import (
	"github.com/apmath-web/currency/Domain"
)

type Repository struct {
	currentRates Domain.ChangeTable
}

func (repository *Repository) SetAllTable(table Domain.ChangeTable) error {
	repository.currentRates = table
	return nil
}

func (repository *Repository) GetAllTable() Domain.ChangeTable {
	return repository.currentRates
}

func (repository *Repository) ClearAllTable() error {
	repository.currentRates = nil
	return nil
}

func (repository *Repository) GetRate(currentCurrency string, wantedCurrency string) Domain.CurrencyRate {
	table := repository.GetAllTable()
	currencies := table.GetCurrencyRate()
	currencyRate := currencies[0]
	for _, cRate := range currencies {
		if cRate.GetCurrentCurrency().GetName() == currentCurrency &&
			cRate.GetWantedCurrency().GetName() == wantedCurrency {
			currencyRate = cRate
			break
		}
	}
	return currencyRate
}
