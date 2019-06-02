package domainModels

import "github.com/apmath-web/currency/Domain"

type CurrencyChange struct {
	amount         int
	baseCurrency   Domain.CurrencyInterface
	wantedCurrency Domain.CurrencyInterface
}

func (i *CurrencyChange) GetBaseCurrency() Domain.CurrencyInterface {
	return i.baseCurrency
}

func (i *CurrencyChange) GetAmount() int {
	return i.amount
}

func (i *CurrencyChange) GetWantedCurrency() Domain.CurrencyInterface {
	return i.wantedCurrency
}

func GenCurrencyChange(baseCurrency Domain.CurrencyInterface, wantedCurrency Domain.CurrencyInterface, amount int) Domain.CurrencyChangeInterface {
	return{amount, baseCurrency, wantedCurrency}
}
