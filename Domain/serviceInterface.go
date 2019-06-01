package Domain

type UpdaterInterface interface {
	Update() error
}

type ExchangerInterface interface {
	Exchange(data CurrencyChangeInterface) float64
}

type FetcherInterface interface {
	FetchAll() RatesInterface
}
