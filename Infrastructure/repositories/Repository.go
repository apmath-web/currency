package repositories

import (
	"sync"

	"github.com/apmath-web/currency/Domain"
	domainModels "github.com/apmath-web/currency/Domain/Models"
)

type Repository struct {
	rates map[string]map[string]float64
}

func (repository *Repository) Set(from Domain.CurrencyInterface, to Domain.CurrencyInterface, value Domain.RateInterface) error {
	repository.rates[from.GetName()][to.GetName()] = value.GetRate()
	return nil
}

func (repository *Repository) Get(from Domain.CurrencyInterface, to Domain.CurrencyInterface) Domain.RateInterface {
	return Domain.GenRate(repository.rates[from.GetName()][to.GetName()])
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
