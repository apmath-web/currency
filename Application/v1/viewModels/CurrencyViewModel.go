package viewModels

import (
	"encoding/json"
	"unicode"

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
	if c.Amount == 0 {
		c.validation.AddMessage(validation.GenMessage("amount", "Is null"))
		return false
	}
	if c.Amount < 0 {
		c.validation.AddMessage(validation.GenMessage("amount", "Is negative"))
		return false
	}
	return true
}

func (c *CurrencyViewModel) validateCurrency(currency string, fieldOfMessage string) bool {
	if len(currency) == 0 {
		c.validation.AddMessage(validation.GenMessage(fieldOfMessage, "404"+currency+"currency is not found"))
		return false
	}

	if len(currency) != 3 {
		c.validation.AddMessage(validation.GenMessage(fieldOfMessage, "The length of "+currency+" is not equal to 3"))
		return false
	}
	for i := 0; i < 3; i++ {
		if !unicode.IsUpper(rune(currency[i])) {
			c.validation.AddMessage(validation.GenMessage(fieldOfMessage, currency+" is not in upper case"))
			return false
		}
	}
	return true
}

func (c *CurrencyViewModel) UnmarshalJSON(b []byte) error {
	tmpModel := JsonCurrency{}
	if err := json.Unmarshal(b, &tmpModel); err != nil {
		return err
	}
	c.JsonCurrency = tmpModel
	return nil
}

func (c *CurrencyViewModel) Validate() bool {
	if c.validateAmount() && c.validateCurrency(c.WantedCurrency, "wantedCurrency") && c.validateCurrency(c.CurrentCurrency, "currentCurrency") {
		return true
	}
	return false
}

func (c *CurrencyViewModel) GetValidation() Domain.ValidationInterface {
	return &c.validation
}
