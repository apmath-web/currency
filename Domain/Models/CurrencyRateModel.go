package domainModels

import (
	"github.com/apmath-web/currency/Domain"
)

type CurrencyRate struct {
	baseCurrency   Domain.CurrencyInterface
	wantedCurrency Domain.CurrencyInterface
	rate           float64
}

func (i *CurrencyRate) GetBaseCurrency() Domain.CurrencyInterface {
	return i.baseCurrency
}

func (i *CurrencyRate) GetWantedCurrency() Domain.CurrencyInterface {
	return i.wantedCurrency
}

func (i *CurrencyRate) GetRate() float64 {
	return i.rate
}

func GenCurrencyRate(bCurrency Domain.CurrencyInterface, wCurrency Domain.CurrencyInterface, rate float64) Domain.CurrencyRateInterface {
	return &CurrencyRate{
		bCurrency,
		wCurrency,
		rate,
	}
}
