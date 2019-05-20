package Domain

type Currency interface {
	GetName() string
}

type CurrencyRate interface {
	GetCurrentCurrency() Currency
	GetWantedCurrency() Currency
	GetRate() int
}

type ChangeTable interface {
	GetCurrencyRate() []CurrencyRate
}
