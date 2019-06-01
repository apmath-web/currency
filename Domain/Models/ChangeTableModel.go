package domainModels

import (
	"github.com/apmath-web/currency/Domain"
)

type ChangeTable struct {
	currencyRates []Domain.CurrencyRateInterface
}

func (i *ChangeTable) GetCurrencyRates() []Domain.CurrencyRateInterface {
	return i.currencyRates
}

func GenChangeTableDomainModel(currencyRates []Domain.CurrencyRateInterface) Domain.ChangeTable {
	return &ChangeTable{
		currencyRates,
	}
}
