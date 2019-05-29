package domainModels

import "github.com/apmath-web/currency/Domain"

type CurrencyChangeModel struct {
	currency Domain.Currency
	amount   float64
}

func (i CurrencyChangeModel) GetCurrency() Domain.Currency {
	return i.currency
}

func (i CurrencyChangeModel) GetAmount() float64 {
	return i.amount
}

func GenCurrencyChangeDomainModel(currency Domain.Currency, amount float64) CurrencyChangeModel {
	return CurrencyChangeModel{currency, amount}

}
