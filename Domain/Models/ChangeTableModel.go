package domainModels

import (
	"github.com/apmath-web/currency/Domain"
)

type ChangeTable struct {
	currencyRates []Domain.CurrencyRate
}

func (i ChangeTable) GetCurrencyRates() []Domain.CurrencyRate {
	return i.currencyRates
}

func GenChangeTableDomainModel(currencyRates []Domain.CurrencyRate) ChangeTable {
	return ChangeTable{
		currencyRates,
	}
}
