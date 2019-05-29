package services

import (
	"github.com/apmath-web/currency/Domain"
)

type CurrencyChangerService struct {
	changeTable Domain.ChangeTableRepositoryInterface
}

func Create(changeTable Domain.ChangeTableRepositoryInterface) *CurrencyChangerService{
	var currencyChangerService = new(CurrencyChangerService)
	currencyChangerService.changeTable = changeTable
	return currencyChangerService
}

func (ccs* CurrencyChangerService) CurrencyChanger(baseCurrency string, wantedCurrency string, amount float64) float64{
	rate := ccs.changeTable.GetRate(baseCurrency, wantedCurrency).GetRate()
	return amount * rate
}