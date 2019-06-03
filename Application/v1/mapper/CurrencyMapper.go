package mapper

import (
	"github.com/apmath-web/currency/Application/v1/viewModels"
	"github.com/apmath-web/currency/Domain"
	domainModels "github.com/apmath-web/currency/Domain/Models"
)

func CurrencyMapper(vm viewModels.CurrencyViewModel) Domain.CurrencyChangeInterface {
	currentCurrency := domainModels.GenCurrency(vm.GetCurrentCurrency())
	wantedCurrency := domainModels.GenCurrency(vm.GetWantedCurrency())
	amount := vm.GetAmount()
	return domainModels.GenCurrencyChange(currentCurrency, wantedCurrency, amount)
}
