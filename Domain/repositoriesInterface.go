package Domain

type RepositoryInterface interface {
	Set(from CurrencyInterface, to CurrencyInterface, value RateInterface) error
	Get(from CurrencyInterface, to CurrencyInterface) RateInterface
	Print()
}
