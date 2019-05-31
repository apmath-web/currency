package services

import (
	"github.com/apmath-web/currency/Domain"
	"github.com/apmath-web/currency/Domain/Models"
)

type CurrencyChangerService struct {
	changeTable Domain.ChangeTableRepositoryInterface
}

func Create(changeTable Domain.ChangeTableRepositoryInterface) Domain.CurrencyChangerInterface {
	service := new(CurrencyChangerService)
	service.changeTable = changeTable
	return service
}

func (ccs *CurrencyChangerService) Change(curCurrency domainModels.CurrencyModel, wanCurrency domainModels.CurrencyModel, Amount float64) float64 {
	rate := ccs.changeTable.GetRate(curCurrency.GetName(), wanCurrency.GetName()).GetRate()
	return Amount * rate
}

