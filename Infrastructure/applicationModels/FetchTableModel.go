package applicationModels

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/apmath-web/currency/Domain"
	"github.com/apmath-web/currency/Infrastructure"
)

type FetcherModel struct{}

func (i FetcherModel) FetchRate(baseCurrency string, wantedCurrency string) float64 {
	url := Infrastructure.ApiUrl + wantedCurrency + "&base=" + baseCurrency

	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	type ratesT struct {
		Rate float64
	}

	type dataT struct {
		Base  string `json:"base"`
		Rates ratesT `json:"rates"`
	}

	data := dataT{
		baseCurrency,
		ratesT{},
	}

	errOfUnmurshal := json.Unmarshal(body, &data)
	if errOfUnmurshal != nil {
		log.Fatal(errOfUnmurshal)
	}

	return data.Rates.Rate
}

func GenFetcherApplicationModel() Domain.Fetcher {
	return FetcherModel{}
}
