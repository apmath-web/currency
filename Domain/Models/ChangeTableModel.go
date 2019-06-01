package domainModels

import (
	"github.com/apmath-web/currency/Domain"
)

type ChangeTable struct {
	currencyRates [6]Domain.CurrencyRate
}

func (i *ChangeTable) GetCurrencyRates() [6]Domain.CurrencyRate {
	return i.currencyRates
}

func GenChangeTableDomainModel(currencyRates [6]Domain.CurrencyRate) *ChangeTable {
	return &ChangeTable{
		currencyRates,
	}
}
