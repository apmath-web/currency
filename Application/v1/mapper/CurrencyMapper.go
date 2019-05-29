package mapper

import (
	"github.com/apmath-web/currency/Application/v1/viewModels"
	"github.com/apmath-web/currency/Domain"
	"github.com/apmath-web/currency/Domain/Models"
)

func CurrencyMapper(vm viewModels.CurrencyViewModel) (Domain.CurrencyChange, Domain.Currency) {
	currentCurrency := domainModels.GenCurrencyDomainModel(vm.GetCurrentCurrency())
	wantedCurrency := domainModels.GenCurrencyDomainModel(vm.GetWantedCurrency())
	return domainModels.GenCurrencyChangeDomainModel(currentCurrency, vm.GetAmount()), wantedCurrency
}
