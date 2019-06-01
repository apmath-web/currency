package viewModels

import (
	"encoding/json"

	"github.com/apmath-web/currency/Application/v1/validation"
	"github.com/apmath-web/currency/Domain"
)

type JsonCurrency struct {
	Amount          int    `json:"amount"`
	CurrentCurrency string `json:"currentCurrency"`
	WantedCurrency  string `json:"wantedCurrency"`
}

type CurrencyViewModel struct {
	JsonCurrency
	validation validation.Validation
}

func (c *CurrencyViewModel) GetAmount() int {
	return c.Amount
}

func (c *CurrencyViewModel) GetCurrentCurrency() string {
	return c.CurrentCurrency
}

func (c *CurrencyViewModel) GetWantedCurrency() string {
	return c.WantedCurrency
}

func (c *CurrencyViewModel) MarshalJSON() (b []byte, e error) {
	return json.Marshal(map[string]interface{}{
		"amount":   c.Amount,
		"currency": c.CurrentCurrency,
	})
}

func (c *CurrencyViewModel) validateAmount() bool {
	if c.Amount < 0 {
		c.validation.AddMessage(validation.GenMessage("amount", "Is negative"))
		return false
	}
	return true
}

//выпилил валидацию валют, так как тперь все обрабатываем вроде

func (c *CurrencyViewModel) UnmarshalJSON(b []byte) error {
	tmpModel := JsonCurrency{}
	if err := json.Unmarshal(b, &tmpModel); err != nil {
		return err
	}
	c.JsonCurrency = tmpModel
	return nil
}

func (c *CurrencyViewModel) Hydrate(model Domain.CurrencyChange) {
	c.Amount = model.GetAmount()
	c.CurrentCurrency = model.GetCurrentCurrency().GetName()
	c.WantedCurrency = model.GetWantedCurrency().GetName()
}

func (c *CurrencyViewModel) Validate() bool {
	if c.validateAmount() && c.validateCurrentCurrency() && c.validateWantedCurrency() {
		return true
	}
	return false
}

func (c *CurrencyViewModel) GetValidation() Domain.ValidationInterface {
	return &c.validation
}
