package Domain

type UpdaterTable interface {
	Update(currentTable ChangeTable) (ChangeTable, error)
}

type CurrencyChanger interface {
	Change(wantToChange CurrencyChange, wantedCurrency Currency) CurrencyChange
}

type ChangeTableServiceInterface interface {
	Add(table ChangeTable) error
	Get() ChangeTable
}
