package domainModels

import "github.com/apmath-web/currency/Domain"

type Rate struct {
	value float64
}

func (i *Rate) GetRate() float64 {
	return i.value
}

func GenRate(value float64) Domain.RateInterface {
	return &Rate{
		value,
	}
}
