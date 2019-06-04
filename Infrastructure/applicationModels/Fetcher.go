package applicationModels

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/apmath-web/currency/Domain"
)

type Fetcher struct{}

func (i *Fetcher) FetchAll() []Domain.FetchRateInterface {
	res, err := http.Get(ApiUrl)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	var data map[string]interface{}

	errOfUnmurshal := json.Unmarshal(body, &data)
	if errOfUnmurshal != nil {
		log.Fatal(errOfUnmurshal)
	}

	m := make(map[string]float64)

	rates := data["Valute"].(map[string]interface{})
	var fetchRates []Domain.FetchRateInterface
	fetchRates = append(fetchRates, GenFetchRate("RUB", "RUB", 1))
	for key, value := range rates {
		infoAboutCurrency := value.(map[string]interface{})
		rate := infoAboutCurrency["Value"].(float64)
		fetchRates = append(fetchRates, GenFetchRate("RUB", key, 1/rate))
		fetchRates = append(fetchRates, GenFetchRate(key, "RUB", rate))
		m[key] = rate
	}

	for j := range m {
		for k := range m {
			fetchRates = append(fetchRates, GenFetchRate(j, k, m[j]/m[k]))
		}
	}
	return fetchRates
}

func GenFetcher() Domain.FetcherInterface {
	return &Fetcher{}
}
