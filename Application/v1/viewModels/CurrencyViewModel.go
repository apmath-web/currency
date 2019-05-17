package viewModels

import "encoding/json"

type CurrencyViewModel struct {
	Amount          int    `json:"amount"`
	CurrentCurrency string `json:"currentCurrency"`
	WantedCurrency  string `json:"wantedCurrency"`
	validation      interface{}
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
	if err := json.Unmarshal(b, &c); err != nil {
		return err
	}
	return nil
}
