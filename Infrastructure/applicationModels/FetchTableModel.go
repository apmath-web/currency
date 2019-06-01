package applicationModels

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/apmath-web/currency/Domain"
)

type FetcherModel struct{}

func (i *FetcherModel) FetchRate(baseCurrency string, wantedCurrency string) (float64, error) {
	url := ApiUrl + "symbols=" + wantedCurrency + "&base=" + baseCurrency

	res, err := http.Get(url)

	if err != nil {
		return -1, err
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	// _, err = os.Stdout.Write(body)

	var data map[string]interface{}

	errOfUnmurshal := json.Unmarshal(body, &data)
	if errOfUnmurshal != nil {
		log.Fatal(errOfUnmurshal)
	}

	rates := data["rates"].(map[string]interface{})
	fmt.Print(rates[wantedCurrency])
	//xccrate := rates[wantedCurrency].(float64)
	return rates[wantedCurrency].(float64), nil
}

func GenFetcherApplicationModel() Domain.Fetcher {
	return &FetcherModel{}
}
