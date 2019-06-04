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
	rates := instance.fetcher.FetchAll()

	for _, currency := range rates {
		instance.repository.Set(
			domainModels.GenCurrency(currency.GetBaseCurrency()),
			domainModels.GenCurrency(currency.GetWantedCurrency()),
			domainModels.GenRate(currency.GetRate()))
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
