package domainModels

import (
	"github.com/apmath-web/currency/Domain"
)

type CurrencyRateModel struct {
	basedCurrency  Domain.Currency
	wantedCurrency Domain.Currency
	Rate           float64
}

func (i *CurrencyRateModel) GetBasedCurrency() Domain.Currency {
	return i.basedCurrency
}

func (i *CurrencyRateModel) GetWantedCurrency() Domain.Currency {
	return i.wantedCurrency
}

func (i *CurrencyRateModel) GetRate() float64 {
	return i.Rate
}

func GenCurrencyRateDomainModel(bCurrency Domain.Currency, wCurrency Domain.Currency, rate float64) *CurrencyRateModel {
	return &CurrencyRateModel{
		bCurrency,
		wCurrency,
		rate,
	}
}
