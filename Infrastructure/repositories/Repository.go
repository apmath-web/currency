package repositories

import (
	"fmt"
	"github.com/apmath-web/currency/Domain"
	domainModels "github.com/apmath-web/currency/Domain/Models"
)

type Repository struct {
	Rates map[string]map[string]float64
}

func (repository *Repository) Set(from Domain.CurrencyInterface, to Domain.CurrencyInterface, value Domain.RateInterface) error {
	//mm, ok :=repository.rates[from.GetName()][to.GetName()] = value.GetRate()
	mm, ok := repository.Rates[from.GetName()]

	if !ok {
		mm = make(map[string]float64)
		repository.Rates[from.GetName()] = mm
	}

	repository.Rates[from.GetName()][to.GetName()] = value.GetRate()
	return nil
}

func (repository *Repository) Get(from Domain.CurrencyInterface, to Domain.CurrencyInterface) Domain.RateInterface {
	return domainModels.GenRate(repository.Rates[from.GetName()][to.GetName()])
}

func (repository *Repository) Print() {

	fmt.Print("=======Print repo==========\n")
	for k1 := range repository.Rates {
		for k2 := range repository.Rates[k1] {
			fmt.Print(k1, " ", k2, " ", repository.Rates[k1][k2], "\n")
		}
	}
}
