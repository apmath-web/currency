package Domain

type CurrencyInterface interface {
	GetName() string
}

type AmountInterface interface {
	GetAmount() int
}

type CurrencyRateInterface interface {
	GetBaseCurrency() CurrencyInterface
	GetWantedCurrency() CurrencyInterface
	GetRate() float64
}

type CurrencyChangeInterface interface {
	GetWantedCurrency() CurrencyInterface
	GetBaseCurrency() CurrencyInterface
	GetAmount() int
}

type RateInterface interface {
	GetRate() float64
}
