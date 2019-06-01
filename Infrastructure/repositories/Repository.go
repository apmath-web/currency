package repositories

import (
	"sync"

	"github.com/apmath-web/currency/Domain"
	domainModels "github.com/apmath-web/currency/Domain/Models"
)

type Repository struct {
	currentRates Domain.ChangeTable
}

func (repository *Repository) SetAllTable(table Domain.ChangeTable) error {
	repository.currentRates = table
	return nil
}

func (repository *Repository) GetAllTable() Domain.ChangeTable {
	return repository.currentRates
}

func (repository *Repository) ClearAllTable() error {
	repository.currentRates = nil
	return nil
}

func (repository *Repository) GetRate(currentCurrency string, wantedCurrency string) Domain.CurrencyRateInterface {
	table := repository.GetAllTable()
	currencies := table.GetCurrencyRates()
	currencyRate := currencies[0]
	for _, cRate := range currencies {
		if cRate.GetBasedCurrency().GetName() == currentCurrency &&
			cRate.GetWantedCurrency().GetName() == wantedCurrency {
			currencyRate = cRate
			break
		}
	}
	return currencyRate
}

var once sync.Once
var repo *Repository

func GenRepository() Domain.ChangeTableRepositoryInterface {

	once.Do(func() {
		repo = &Repository{
			domainModels.GenChangeTableDomainModel(
				[]Domain.CurrencyRateInterface{
					domainModels.GenCurrencyRateDomainModel(
						domainModels.GenCurrencyDomainModel("USD"),
						domainModels.GenCurrencyDomainModel("EUR"),
						0),
					domainModels.GenCurrencyRateDomainModel(
						domainModels.GenCurrencyDomainModel("USD"),
						domainModels.GenCurrencyDomainModel("RUB"),
						0),
					domainModels.GenCurrencyRateDomainModel(
						domainModels.GenCurrencyDomainModel("EUR"),
						domainModels.GenCurrencyDomainModel("USD"),
						0),
					domainModels.GenCurrencyRateDomainModel(
						domainModels.GenCurrencyDomainModel("EUR"),
						domainModels.GenCurrencyDomainModel("RUB"),
						0),
					domainModels.GenCurrencyRateDomainModel(
						domainModels.GenCurrencyDomainModel("RUB"),
						domainModels.GenCurrencyDomainModel("EUR"),
						0),
					domainModels.GenCurrencyRateDomainModel(
						domainModels.GenCurrencyDomainModel("RUB"),
						domainModels.GenCurrencyDomainModel("USD"),
						0),
				}),
		}
	})

	return repo
}
