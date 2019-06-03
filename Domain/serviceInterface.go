package Domain

type UpdaterTable interface {
	Update(currentTable ChangeTable) (ChangeTable, error)
}

type CurrencyChangerInterface interface {
	Change(curCurrency Currency, wanCurrency Currency, Amount float64) float64
}

type FetchRateInterface interface {
	GetCurrency() string
	GetRate() float64
}

type FetcherInterface interface {
	FetchAll() []FetchRateInterface
}
