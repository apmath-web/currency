package domainModels

import "github.com/apmath-web/currency/Domain"

type CurrencyChangeModel struct {
	amount          int
	currentCurrency Domain.Currency
	wantedCurrency  Domain.Currency
}

func (i *CurrencyChangeModel) GetCurrentCurrency() Domain.Currency {
	return i.currentCurrency
}

func (i *CurrencyChangeModel) GetAmount() int {
	return i.amount
}

func (i *CurrencyChangeModel) GetWantedCurrency() Domain.Currency {
	return i.wantedCurrency
}

func GenCurrencyChangeDomainModel(currentCurrency Domain.Currency, wantedCurrency Domain.Currency, amount int) CurrencyChangeModel {
	return CurrencyChangeModel{amount, currentCurrency, wantedCurrency}
}
