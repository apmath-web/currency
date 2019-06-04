package services

import (
	"github.com/apmath-web/currency/Domain"
	domainModels "github.com/apmath-web/currency/Domain/Models"
)

type Exchanger struct {
	repository Domain.RepositoryInterface
}

func (instance *Exchanger) Exchange(data Domain.CurrencyChangeInterface) (Domain.AmountInterface, error) {
	rate, err := instance.repository.Get(domainModels.GenCurrency(data.GetWantedCurrency()), domainModels.GenCurrency(data.GetBaseCurrency()))
	if err != nil {
		return nil, err
	}
	return domainModels.GenAmount(int(float64(data.GetAmount()) * rate.GetRate())), nil
}

func GenExchanger(repository Domain.RepositoryInterface) Domain.ExchangerInterface {
	return &Exchanger{
		repository,
	}
}
