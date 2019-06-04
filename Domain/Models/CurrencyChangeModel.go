package domainModels

import "github.com/apmath-web/currency/Domain"

type CurrencyChange struct {
	amount         int
	baseCurrency   string
	wantedCurrency string
}

func (i *CurrencyChange) GetBaseCurrency() string {
	return i.baseCurrency
}

func (i *CurrencyChange) GetAmount() int {
	return i.amount
}

func (i *CurrencyChange) GetWantedCurrency() string {
	return i.wantedCurrency
}

func GenCurrencyChange(baseCurrency string, wantedCurrency string, amount int) Domain.CurrencyChangeInterface {
	return &CurrencyChange{amount, baseCurrency, wantedCurrency}
}
