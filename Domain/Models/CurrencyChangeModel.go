package domainModels

import "github.com/apmath-web/currency/Domain"

type CurrencyChangeModel struct {
	currency Domain.Currency
	amount   int
}

func (i CurrencyChangeModel) GetCurrency() Domain.Currency {
	return i.currency
}

func (i CurrencyChangeModel) GetAmount() int {
	return i.amount
}

func GenCurrencyChangeDomainModel(currency Domain.Currency, amount int) CurrencyChangeModel {
	return CurrencyChangeModel{currency, amount}

}
