package applicationModels

import (
	"github.com/apmath-web/currency/Domain"
)

type Fetcher struct{}

	body, err := ioutil.ReadAll(res.Body)

	var data map[string]interface{}

	errOfUnmurshal := json.Unmarshal(body, &data)
	if errOfUnmurshal != nil {
		log.Fatal(errOfUnmurshal)
	}

	rates := data["rates"].(map[string]interface{})
	fmt.Print(rates[wantedCurrency])
	return rates[wantedCurrency].(float64), nil
}

}
