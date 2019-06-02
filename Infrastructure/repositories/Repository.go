package repositories

import (
	"github.com/apmath-web/currency/Domain"
)

type Repository struct {
	currentRates Domain.RatesInterface
}

func (repository *Repository) SetAllTable(table Domain.RatesInterface) error {
	repository.currentRates = table
	return nil
}

func (repository *Repository) GetAllTable() Domain.RatesInterface {
	return repository.currentRates
}

func (repository *Repository) ClearAllTable() error {
	repository.currentRates = nil
	return nil
}

func (repository *Repository) GetRate(currentCurrency string, wantedCurrency string) Domain.CurrencyRateInterface {
	table := repository.GetAllTable()
	currencies := table.GetCurrencyRates()
	currencyRate := currencies[0]
	for _, cRate := range currencies {
		if cRate.GetBasedCurrency().GetName() == currentCurrency &&
			cRate.GetWantedCurrency().GetName() == wantedCurrency {
			currencyRate = cRate
			break
		}
	}
	return currencyRate
}
