package Domain

type UpdaterTable interface {
	Update(currentTable ChangeTable) (ChangeTable, error)
}

type CurrencyChangerInterface interface {
	Change(curCurrency Currency, wanCurrency Currency, Amount float64) float64
}

type ChangeTableServiceInterface interface {
	Add(table ChangeTable) error
	Get() ChangeTable
}
