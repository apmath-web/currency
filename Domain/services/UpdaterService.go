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
	for _, currency1 := range rates {
		for _, currency2 := range rates {
			_ = instance.repository.Set(domainModels.GenCurrency(currency1.GetCurrency()),
				domainModels.GenCurrency(currency2.GetCurrency()),
				domainModels.GenRate(1/(1/currency1.GetRate())/(1/currency2.GetRate())))
		}
	}
}
