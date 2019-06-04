package viewModels

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCreditViewModel(t *testing.T) {
	js := map[string]interface{}{
		"amount":          200,
		"currentCurrency": "USD",
		"wantedCurrency":  "EUR",
	}
	strTest, _ := json.Marshal(js)
	fmt.Println(string(strTest))
	m := new(CurrencyViewModel)
	_ = json.Unmarshal(strTest, m)
	fmt.Printf("%+v\n", m)
	strResult, _ := json.Marshal(m)
	fmt.Println(string(strResult))

}
