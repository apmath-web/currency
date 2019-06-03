package services

import (
	"github.com/apmath-web/currency/Domain"
	domainModels "github.com/apmath-web/currency/Domain/Models"
)

type Updater struct {
	repository Domain.RepositoryInterface
	fetcher    Domain.FetcherInterface
}

func (instance *Updater) Update() error {
	instance.repository.Set(
		domainModels.GenCurrency("RUB"),
		domainModels.GenCurrency("RUB"),
		domainModels.GenRate(1.0))

	rates := instance.fetcher.FetchAll()

	for _, currency := range rates {
		instance.repository.Set(
			domainModels.GenCurrency("RUB"),
			domainModels.GenCurrency(currency.GetCurrency()),
			domainModels.GenRate(currency.GetRate()),
		)
		instance.repository.Set(
			domainModels.GenCurrency(currency.GetCurrency()),
			domainModels.GenCurrency("RUB"),
			domainModels.GenRate(1/currency.GetRate()),
		)

	}

	for _, currency1 := range rates {
		for _, currency2 := range rates {
			instance.repository.Set(
				domainModels.GenCurrency(currency1.GetCurrency()),
				domainModels.GenCurrency(currency2.GetCurrency()),
				domainModels.GenRate(currency1.GetRate()/currency2.GetRate()))
		}
	}

	// instance.repository.Print()
	return nil
}

func GenUpdater(repository Domain.RepositoryInterface, fetcher Domain.FetcherInterface) Domain.UpdaterInterface {
	return &Updater{
		repository,
		fetcher,
	}
}
