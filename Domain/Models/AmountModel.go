package domainModels

import "github.com/apmath-web/currency/Domain"

type Amount struct {
	value int
}

func (i *Amount) GetAmount() int {
	return i.value
}

func GenAmount(value int) Domain.AmountInterface {
	return &Amount{
		value,
	}
}
