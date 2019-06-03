package services

import (
	"github.com/apmath-web/currency/Domain"
	domainModels "github.com/apmath-web/currency/Domain/Models"
)

type Exchanger struct {
	repository Domain.RepositoryInterface
}

func (instance *Exchanger) Exchange(data Domain.CurrencyChangeInterface) Domain.AmountInterface {
	rate := instance.repository.Get(data.GetWantedCurrency(), data.GetBaseCurrency())
	return domainModels.GenAmount(int(float64(data.GetAmount()) * rate.GetRate()))
}

func GenExchanger(repository Domain.RepositoryInterface) Domain.ExchangerInterface {
	return &Exchanger{
		repository,
	}
}
