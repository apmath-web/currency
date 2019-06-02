package domainModels

import (
	"github.com/apmath-web/currency/Domain"
)

type Rates struct {
	currencyRates []Domain.CurrencyRateInterface
}

func (i *Rates) GetRates() []Domain.CurrencyRateInterface {
	return i.currencyRates
}

func GenRates(currencyRates []Domain.CurrencyRateInterface) Domain.RatesInterface {
	return Rates{
		currencyRates,
	}
}
