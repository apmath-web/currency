package Domain

type RepositoryInterface interface {
	Set(data CurrencyRateInterface) error
	Get(from string, to string) RateInterface
}
