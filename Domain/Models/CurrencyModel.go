package domainModels

import "github.com/apmath-web/currency/Domain"

type Currency struct {
	name string
}

func (i *Currency) GetName() string {
	return i.name
}

func GenCurrency(name string) Domain.CurrencyInterface {
	return Currency{
		name,
	}
}
