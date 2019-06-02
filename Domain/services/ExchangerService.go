package services

import (
	"github.com/apmath-web/currency/Domain"
)

type Exchanger struct {
	rates     Domain.RepositoryInterface
	exchanger Domain.ExchangerInterface
}

func (exchanger *Exchanger) Exchange(data Domain.CurrencyChangeInterface) Domain.AmountInterface {
	//rate := exchanger.changeTable.GetRate(curCurrency.GetName(), wanCurrency.GetName()).GetRate()
	rate:=rates
	return data.GetAmount() * rate
}
