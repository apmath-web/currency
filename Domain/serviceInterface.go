package Domain

type UpdaterInterface interface {
	Update() error
}

type ExchangerInterface interface {
	Exchange(data CurrencyChangeInterface) AmountInterface
}

type FetchRateInterface interface {
	GetWantedCurrency() string
	GetBaseCurrency() string
	GetRate() float64
}

type FetcherInterface interface {
	FetchAll() []FetchRateInterface
}
