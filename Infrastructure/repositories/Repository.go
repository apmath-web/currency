package repositories

import (
	"log"
	"sync"

	"github.com/apmath-web/currency/Domain"
	domainModels "github.com/apmath-web/currency/Domain/Models"
)

type Repository struct {
	rates map[string]map[string]float64
}

func (repository *Repository) Set(from Domain.CurrencyInterface, to Domain.CurrencyInterface, value Domain.RateInterface) error {
	mm, ok := repository.rates[from.GetName()]

	if !ok {
		mm = make(map[string]float64)
		repository.rates[from.GetName()] = mm
	}

	repository.rates[from.GetName()][to.GetName()] = value.GetRate()
	return nil
}

func (repository *Repository) Get(from Domain.CurrencyInterface, to Domain.CurrencyInterface) Domain.RateInterface {

	_, ok := repository.rates[from.GetName()]
	if !ok {
		log.Fatal("Currency" + from.GetName() + " doesn't exist")
	}
	val2, ok2 := repository.rates[from.GetName()][to.GetName()]

	if !ok2 {
		log.Fatal("Currency " + to.GetName() + " doesn't exist")
	}
	return domainModels.GenRate(val2)
}

var repo *Repository
var once sync.Once

func GenRepository() Domain.RepositoryInterface {
	once.Do(func() {
		repo = &Repository{make(map[string]map[string]float64)}
	})
	return repo
}
