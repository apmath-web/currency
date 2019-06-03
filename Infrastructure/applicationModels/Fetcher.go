package applicationModels

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/apmath-web/currency/Domain"
)

type Fetcher struct{}

func (i *Fetcher) FetchAll() []Domain.FetchRateInterface {
	res, err := http.Get(ApiUrl)

	if err != nil {
		return nil
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	var data map[string]interface{}

	errOfUnmurshal := json.Unmarshal(body, &data)
	if errOfUnmurshal != nil {
		log.Fatal(errOfUnmurshal)
	}

	rates := data["Valute"].(map[string]interface{})
	var fetchRates []Domain.FetchRateInterface
	for key, value := range rates {
		infoAboutCurrency := value.(map[string]interface{})
		rate := infoAboutCurrency["Value"].(float64)
		fetchRates = append(fetchRates, GenFetchRate(key, rate))
	}
	for _, obj := range fetchRates {
		fmt.Print(obj.GetCurrency(), " ", obj.GetRate(), "\n")
	}
	return fetchRates

}

func GenFetcher() Domain.FetcherInterface {
	return &Fetcher{}
}
