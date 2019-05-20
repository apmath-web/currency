package Domain

type ChangeTableRepositoryInterface interface {
	SetAllTable(table ChangeTable) error
	GetAllTable() ChangeTable
	ClearAllTable() error
}
