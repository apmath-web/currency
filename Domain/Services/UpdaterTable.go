package services

import (
	"github.com/apmath-web/currency/Domain"
	"github.com/apmath-web/currency/Domain/domainModels"

	infrastructure "github.com/apmath-web/currency/Infrastructure"
)

type UpdaterTableModel struct {
	fetcher infrastructure.Fetcher
}

func (i UpdaterTableModel) Update(curTable Domain.ChangeTable) (Domain.ChangeTable, error) {
	curRates := curTable.GetCurrencyRates()
	for ind, rate := range curRates {
		wantedCurrency := rate.GetWantedCurrency()
		baseCurrency := rate.GetBasedCurrency()
		wantedCurrencyName := wantedCurrency.GetName()
		baseCurrencyName := baseCurrency.GetName()
		actualeRate := i.fetcher.FetchRate(baseCurrencyName, wantedCurrencyName)

		curRates[ind] = domainModels.GenCurrencyRateDomainModel(baseCurrency, wantedCurrency, actualeRate)
	}
}

func GenUpdaterTableDomainModel() Domain.UpdaterTable {
	return UpdaterTableModel{}
}
