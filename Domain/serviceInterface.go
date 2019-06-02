package Domain

type UpdaterInterface interface {
	Update() error
}

type ExchangerInterface interface {
	Exchange(data CurrencyChangeInterface) AmountInterface
}

type FetcherInterface interface {
	FetchAll() RatesInterface
}
