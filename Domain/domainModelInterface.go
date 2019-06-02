package Domain

type CurrencyInterface interface {
	GetName() string
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

type RatesInterface interface {
	GetCurrencyRates() []CurrencyRateInterface
}

type RateInterface interface {
	GetRate() float64
}
