package Domain

type Currency interface {
	GetName() string
}

type CurrencyRate interface {
	GetCurrentCurrency() Currency
	GetWantedCurrency() Currency
	GetRate() float64
}

type CurrencyChange interface {
	GetCurrency() Currency
	GetAmount() int
}

type ChangeTable interface {
	GetCurrencyRates() []CurrencyRate
}
