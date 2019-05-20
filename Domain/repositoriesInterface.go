package Domain

type CurrencyRateRepositoryInterface interface {
	GetCurrencyRate(curCurrency Currency, wanCurrency Currency) CurrencyRate
	SetCurrencyRate(rate CurrencyRate) error
	UpdateCurrencyRate(rate CurrencyRate) error
}
