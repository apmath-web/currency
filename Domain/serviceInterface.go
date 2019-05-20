package Domain

type UpdaterTable interface {
	Update() (ChangeTable, error)
}

type CurrencyChanger interface {
	Change(curCurrency Currency, wanCurrency Currency, Amount int) CurrencyChange
}

type ChangeTableServiceInterface interface {
	Add(table ChangeTable) error
	Get() ChangeTable
}
