package Domain

type UpdaterTable interface {
	Update() (ChangeTable, error)
}

type CurrencyChanger interface {
	Change(curCurrency Currency, wanCurrency Currency) CurrencyChange
}

type ChangeTableServiceInterface interface {
	Add(table ChangeTable) error
	Get() ChangeTable
}
