package applicationModels

import (
	"github.com/apmath-web/currency/Domain"
)

type FetchRate struct {
	wantedCurrency string
	baseCurrency   string
	rate           float64
}

func (i *FetchRate) GetWantedCurrency() string {
	return i.wantedCurrency
}

func (i *FetchRate) GetBaseCurrency() string {
	return i.baseCurrency
}

func (i *FetchRate) GetRate() float64 {
	return i.rate
}

func GenFetchRate(wantedCurrency string, baseCurrency string, rate float64) Domain.FetchRateInterface {
	return &FetchRate{
		wantedCurrency,
		baseCurrency,
		rate,
	}
}
