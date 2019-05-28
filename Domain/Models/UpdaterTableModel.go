package domainModels

import (
	"github.com/apmath-web/currency/Domain"
	infrastructure "github.com/apmath-web/currency/Infrastructure"
)

type UpdaterTableModel struct {
	fetcher infrastructure.Fetcher
}

func (i UpdaterTableModel) Update(curTable Domain.ChangeTable) (Domain.ChangeTable, error) {

}

func GenUpdaterTableDomainModel() Domain.UpdaterTable {
	return UpdaterTableModel{}
}
