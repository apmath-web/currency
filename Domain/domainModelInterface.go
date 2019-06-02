package Domain

type Currency interface {
	GetName() string
}

type AmountInterface interface {
	GetAmount() int
}

type CurrencyRateInterface interface {
	GetBasedCurrency() Currency
	GetWantedCurrency() Currency
	GetRate() float64
}

type CurrencyChangeInterface interface {
	GetWantedCurrency() Currency
	GetCurrentCurrency() Currency
	GetAmount() int
}

type RatesInterface interface {
	GetCurrencyRates() []CurrencyRateInterface
}

type RateInterface interface {
	GetRate() float64
}
