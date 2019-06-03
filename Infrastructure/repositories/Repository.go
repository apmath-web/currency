package repositories

import (
	"fmt"
	"sync"

	"github.com/apmath-web/currency/Domain"
	domainModels "github.com/apmath-web/currency/Domain/Models"
)

type Repository struct {
	rates map[string]map[string]float64
}

func (repository *Repository) Set(from Domain.CurrencyInterface, to Domain.CurrencyInterface, value Domain.RateInterface) error {
	//mm, ok :=repository.rates[from.GetName()][to.GetName()] = value.GetRate()
	mm, ok := repository.rates[from.GetName()]

	if !ok {
		mm = make(map[string]float64)
		repository.rates[from.GetName()] = mm
	}

	repository.rates[from.GetName()][to.GetName()] = value.GetRate()
	return nil
}

func (repository *Repository) Get(from Domain.CurrencyInterface, to Domain.CurrencyInterface) Domain.RateInterface {
	return domainModels.GenRate(repository.rates[from.GetName()][to.GetName()])
}

func (repository *Repository) Print() {

	fmt.Print("=======Print repo==========\n")
	for k1 := range repository.rates {
		for k2 := range repository.rates[k1] {
			fmt.Print(k1, " ", k2, " ", repository.rates[k1][k2], "\n")
		}
	}
}

var repo *Repository
var once sync.Once

func GenRepository() Domain.RepositoryInterface {
	once.Do(func() {
		repo = &Repository{make(map[string]map[string]float64)}
	})
	return repo
}
