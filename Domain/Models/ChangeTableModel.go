package domainModels

import (
	"github.com/apmath-web/currency/Domain"
)

type ChangeTable struct {
	currencyRates [6]Domain.CurrencyRateInterface
}

func (i *ChangeTable) GetCurrencyRates() [6]Domain.CurrencyRateInterface {
	return i.currencyRates
}

func GenChangeTableDomainModel(currencyRates [6]Domain.CurrencyRateInterface) Domain.ChangeTable {
	return &ChangeTable{
		currencyRates,
	}
}
