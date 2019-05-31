package Domain

import domainModels "github.com/apmath-web/currency/Domain/Models"

type UpdaterTable interface {
	Update(currentTable ChangeTable) (ChangeTable, error)
}

type CurrencyChangerInterface interface {
	Change(curCurrency domainModels.CurrencyModel, wanCurrency domainModels.CurrencyModel, Amount float64) float64
}

type ChangeTableServiceInterface interface {
	Add(table ChangeTable) error
	Get() ChangeTable
}
