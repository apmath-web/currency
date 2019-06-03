package applicationModels

import (
	"github.com/apmath-web/currency/Domain"
)

type FetchRate struct {
	currency string
	rate     float64
}

func (i *FetchRate) GetCurrency() string {
	return i.currency
}

func (i *FetchRate) GetRate() float64 {
	return i.rate
}

func GenFetchRate(currency string, rate float64) Domain.FetchRateInterface {
	return &FetchRate{
		currency,
		rate,
	}
}
