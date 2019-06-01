package Domain

type Currency interface {
	GetName() string
}

type CurrencyRateInterface interface {
	GetBasedCurrency() Currency
	GetWantedCurrency() Currency
	GetRate() float64
}

type CurrencyChange interface {
	GetWantedCurrency() Currency
	GetCurrentCurrency() Currency
	GetAmount() int
}

type ChangeTable interface {
	GetCurrencyRates() []CurrencyRateInterface
}

type Fetcher interface {
	FetchRate(baseCurrency string, wantedCurrency string) (float64, error)
}
