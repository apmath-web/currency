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
	GetAmount() int
}

type ChangeTable interface {
	GetCurrencyRates() []CurrencyRate
}

type Fetcher interface {
	FetchRate(baseCurrency string, wantedCurrency string) (float64, error)
}
