package repositories

import (
	"github.com/apmath-web/currency/Domain"
	"log"
	"sync"
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

var repo *Repository
var once sync.Once

func GenRepository() Domain.RepositoryInterface {
	once.Do(func() {
		repo = &Repository{make(map[string]map[string]float64)}
	})
	return repo
}
