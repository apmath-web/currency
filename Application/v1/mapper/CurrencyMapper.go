package mapper

import (
	"github.com/apmath-web/currency/Application/v1/viewModels"
	domainModels "github.com/apmath-web/currency/Domain/Models"
)

func CurrencyMapper(vm viewModels.CurrencyViewModel) domainModels.CurrencyChangeModel {
	currentCurrency := domainModels.GenCurrencyDomainModel(vm.GetCurrentCurrency())
	wantedCurrency := domainModels.GenCurrencyDomainModel(vm.GetWantedCurrency())
	amount := vm.GetAmount()
	return domainModels.GenCurrencyChangeDomainModel(currentCurrency, wantedCurrency, amount)
}
