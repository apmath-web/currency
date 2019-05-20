package viewModels

import (
	"encoding/json"
	"github.com/apmath-web/currency/Domain"
)

type JsonCurrency struct {
	Amount          int    `json:"amount"`
	CurrentCurrency string `json:"currentCurrency"`
	WantedCurrency  string `json:"wantedCurrency"`
}

type CurrencyViewModel struct {
	JsonCurrency
	// later use hear realisation not interface
	validation Domain.ValidationInterface
}

func (c *CurrencyViewModel) GetAmount() int {
	return c.Amount
}

func (c *CurrencyViewModel) GetCurrentCurrency() int {
	return c.Amount
}

func (c *CurrencyViewModel) GetWantedCurrency() int {
	return c.Amount
}

func (c *CurrencyViewModel) MarshalJSON() (b []byte, e error) {
	return json.Marshal(map[string]interface{}{
		"amount":   c.Amount,
		"currency": c.CurrentCurrency,
	})
}

func (c *CurrencyViewModel) UnmarshalJSON(b []byte) error {
	tmpModel := JsonCurrency{}
	if err := json.Unmarshal(b, &tmpModel); err != nil {
		return err
	}
	c.JsonCurrency = tmpModel
	return nil
}

func (c *CurrencyViewModel) Hydrate(model Domain.CurrencyChange) {

}

func (c *CurrencyViewModel) Validate() bool {
	return true
}

func (c *CurrencyViewModel) GetValidation() Domain.ValidationInterface {
	return c.validation
}
