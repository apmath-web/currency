package domainModels

import (
	"github.com/apmath-web/currency/Domain"
)

type CurrencyRateModel struct {
	currentCurrency Domain.Currency
	wantedCurrency  Domain.Currency
}

func (i CurrencyRateModel) GetCurrentCurrency() Domain.Currency {
	return i.currentCurrency
}

func (i CurrencyRateModel) GetWantedCurrency() Domain.Currency {
	return i.wantedCurrency
}
