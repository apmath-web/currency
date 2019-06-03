package Domain

type UpdaterInterface interface {
	Update() error
}

type ExchangerInterface interface {
	Exchange(data CurrencyChangeInterface) AmountInterface
}

type FetchRateInterface interface {
	GetBaseCurrency() string
	GetWantedCurrency() string
	GetRate() float64
}

type FetcherInterface interface {
	FetchAll() []FetchRateInterface
}
