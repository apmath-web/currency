package services

import (
	"log"

	"github.com/apmath-web/currency/Domain"
	domainModels "github.com/apmath-web/currency/Domain/Models"
)

type UpdaterTableModel struct {
	fetcher Domain.Fetcher
}

func (i *UpdaterTableModel) Update(currentTable Domain.ChangeTable) (Domain.ChangeTable, error) {
	currentRates := currentTable.GetCurrencyRates()
	for ind, rate := range currentRates {
		wantedCurrency := rate.GetWantedCurrency()
		baseCurrency := rate.GetBasedCurrency()
		wantedCurrencyName := wantedCurrency.GetName()
		baseCurrencyName := baseCurrency.GetName()
		actualeRate, err := i.fetcher.FetchRate(baseCurrencyName, wantedCurrencyName)

		if err != nil {
			log.Fatal(err)
		}
		currentRates[ind] = domainModels.GenCurrencyRateDomainModel(baseCurrency, wantedCurrency, actualeRate)
	}

	return domainModels.GenChangeTableDomainModel(currentRates), nil
}

func GenUpdaterTableDomainModel(fetcher Domain.Fetcher) Domain.UpdaterTable {
	return &UpdaterTableModel{
		fetcher,
	}
}
