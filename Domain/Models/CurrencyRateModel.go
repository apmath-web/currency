package domainModels

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/apmath-web/currency/Domain"
)

type CurrencyRateModel struct {
	basedCurrency  Domain.Currency
	wantedCurrency Domain.Currency
}

func (i CurrencyRateModel) GetBasedCurrency() Domain.Currency {
	return i.basedCurrency
}

func (i CurrencyRateModel) GetWantedCurrency() Domain.Currency {
	return i.wantedCurrency
}

func (i CurrencyRateModel) GetRate() float64 {
	url := "https://api.exchangeratesapi.io/latest?symbols=" + i.wantedCurrency.GetName() + "&base=" + i.basedCurrency.GetName()
	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	type ratesT struct {
		rate float64
	}

	type dataT struct {
		Base  string `json:"base"`
		Rates ratesT `json:"rates`
	}

	data := dataT{
		i.basedCurrency.GetName(),
		ratesT{},
	}

	errOfUnmurshal := json.Unmarshal(body, &data)
	if errOfUnmurshal != nil {
		log.Fatal(errOfUnmurshal)
	}

	return data.Rates.rate

}
