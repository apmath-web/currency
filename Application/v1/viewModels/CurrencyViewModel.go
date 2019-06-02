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
	if c.Amount < 0 {
		c.validation.AddMessage(validation.GenMessage("amount", "Is negative"))
		return false
	}
	return true
}

func (c *CurrencyViewModel) validateCurrentCurrency() bool {
	if len(c.CurrentCurrency) != 3 {
		c.validation.AddMessage(validation.GenMessage("CurrentCurrency", "Is incorrect"))
		return false
	}

	for _, char := range c.CurrentCurrency {
		if !unicode.IsUpper(char) {
			c.validation.AddMessage(validation.GenMessage("CurrentCurrency", "Is incorrect"))
			return false
		}
	}
	return true
}

func (c *CurrencyViewModel) validateWantedCurrency() bool {
	if len(c.WantedCurrency) != 3 {
		c.validation.AddMessage(validation.GenMessage("WantedCurrency", "Is incorrect"))
		return false
	}

	for _, char := range c.WantedCurrency {
		if !unicode.IsUpper(char) {
			c.validation.AddMessage(validation.GenMessage("WantedCurrency", "Is incorrect"))
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

func (c *CurrencyViewModel) Hydrate(model Domain.CurrencyChangeInterface) {
	c.Amount = model.GetAmount()
	c.CurrentCurrency = model.GetBaseCurrency().GetName()
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
