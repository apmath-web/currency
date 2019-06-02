package Domain

type RepositoryInterface interface {
	Set(data CurrencyRateInterface) error
	Get(from Currency, to Currency) RateInterface
}
