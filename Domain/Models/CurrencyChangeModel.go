package domainModels

import "github.com/apmath-web/currency/Domain"

type CurrencyChange struct {
	amount         int
<<<<<<< HEAD
	baseCurrency   string
	wantedCurrency string
}

func (i *CurrencyChange) GetBaseCurrency() string {
=======
	baseCurrency   Domain.CurrencyInterface
	wantedCurrency Domain.CurrencyInterface
}

func (i *CurrencyChange) GetBaseCurrency() Domain.CurrencyInterface {
>>>>>>> dev
	return i.baseCurrency
}

func (i *CurrencyChange) GetAmount() int {
	return i.amount
}

<<<<<<< HEAD
func (i *CurrencyChange) GetWantedCurrency() string {
	return i.wantedCurrency
}

func GenCurrencyChange(baseCurrency string, wantedCurrency string, amount int) Domain.CurrencyChangeInterface {
=======
func (i *CurrencyChange) GetWantedCurrency() Domain.CurrencyInterface {
	return i.wantedCurrency
}

func GenCurrencyChange(baseCurrency Domain.CurrencyInterface, wantedCurrency Domain.CurrencyInterface, amount int) Domain.CurrencyChangeInterface {
>>>>>>> dev
	return &CurrencyChange{amount, baseCurrency, wantedCurrency}
}
