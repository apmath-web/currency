package Services

import (
	"log"

	"github.com/apmath-web/currency/Domain"
	domainModels "github.com/apmath-web/currency/Domain/Models"
	"github.com/apmath-web/currency/Infrastructure/applicationModels"
)

type UpdaterTableModel struct {
	fetcher Domain.Fetcher
}

func (i *UpdaterTableModel) Update(curTable Domain.ChangeTable) (Domain.ChangeTable, error) {
	curRates := curTable.GetCurrencyRates()
	for ind, rate := range curRates {
		wantedCurrency := rate.GetWantedCurrency()
		baseCurrency := rate.GetBasedCurrency()
		wantedCurrencyName := wantedCurrency.GetName()
		baseCurrencyName := baseCurrency.GetName()
		actualeRate, err := i.fetcher.FetchRate(baseCurrencyName, wantedCurrencyName)

		if err != nil {
			log.Fatal(err)
		}
		curRates[ind] = domainModels.GenCurrencyRateDomainModel(baseCurrency, wantedCurrency, actualeRate)
	}

	return domainModels.GenChangeTableDomainModel(curRates), nil
}

func GenUpdaterTableDomainModel() Domain.UpdaterTable {
	return &UpdaterTableModel{
		applicationModels.GenFetcherApplicationModel(),
	}
}
