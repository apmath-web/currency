package domainModels

import (
	"github.com/apmath-web/currency/Domain"
)

type CurrencyRateModel struct {
	basedCurrency  Domain.Currency
	wantedCurrency Domain.Currency
	Rate           int
}

func (i *CurrencyRateModel) GetBasedCurrency() Domain.Currency {
	return i.basedCurrency
}

func (i *CurrencyRateModel) GetWantedCurrency() Domain.Currency {
	return i.wantedCurrency
}

func (i *CurrencyRateModel) GetRate() int {
	return i.Rate
}

func GenCurrencyRateDomainModel(bCurrency Domain.Currency, wCurrency Domain.Currency, rate int) CurrencyRateModel {
	return CurrencyRateModel{
		bCurrency,
		wCurrency,
		rate,
	}
}
