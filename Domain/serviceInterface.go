package Domain

type UpdaterInterface interface {
	Update() error
}

type ExchangerInterface interface {
	Exchange(data CurrencyChangeInterface) AmountInterface
<<<<<<< HEAD
}

type FetchRateInterface interface {
	GetBaseCurrency() string
	GetWantedCurrency() string
	GetRate() float64
=======
>>>>>>> changes ExchangeInterface type of what you get from Exchange
}

type FetcherInterface interface {
	FetchAll() []FetchRateInterface
}
