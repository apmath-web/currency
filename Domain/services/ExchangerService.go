package services

import (
	"github.com/apmath-web/currency/Domain"
	"github.com/apmath-web/currency/Domain/Models"
)

type Exchanger struct {
	repository Domain.RepositoryInterface
}

func (instance *Exchanger) Exchange(data Domain.CurrencyChangeInterface) Domain.AmountInterface {
	rate := instance.repository.Get(data.GetWantedCurrency(), data.GetBaseCurrency())
	return domainModels.GenAmount(int(float64(data.GetAmount()) * rate.GetRate()))
}
