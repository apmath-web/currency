package repositories

import (
	"errors"
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

func (repository *Repository) Get(from Domain.CurrencyInterface, to Domain.CurrencyInterface) (Domain.RateInterface, error) {

	_, ok := repository.rates[from.GetName()]
	if !ok {
		return nil, errors.New("Currency " + from.GetName() + " doesn't exist")
	}
	val2, ok2 := repository.rates[from.GetName()][to.GetName()]

	if !ok2 {
		return nil, errors.New("Currency " + to.GetName() + " doesn't exist")
	}
	return domainModels.GenRate(val2), nil
}

func (repository *Repository) IsEmpty() bool {
	if len(repository.rates) == 0 {
		return true
	}
	return false
}

var repo *Repository
var once sync.Once

func GenRepository() Domain.RepositoryInterface {
	once.Do(func() {
		repo = &Repository{make(map[string]map[string]float64)}
	})
	return repo
}
