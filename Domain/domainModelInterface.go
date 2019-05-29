package Domain

type Currency interface {
	GetName() string
}

type CurrencyRate interface {
	GetBasedCurrency() Currency
	GetWantedCurrency() Currency
	GetRate() float64
}

type CurrencyChange interface {
	GetCurrency() Currency
	GetAmount() float64
}

type ChangeTable interface {
	GetCurrencyRates() []CurrencyRate
}
