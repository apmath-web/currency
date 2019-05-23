package services

import (
	"github.com/apmath-web/currency/Domain"
	"github.com/apmath-web/currency/Infrastructure/repositories"
)

type UpdaterTableService struct {
	repositiry repositories.Repository
}

func (updater UpdaterTableService) Update() (Domain.ChangeTable, error) {
	var rates = updater.repositiry.GetAllTable().GetCurrencyRates()

	return nil, nil
}
