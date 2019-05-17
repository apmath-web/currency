package viewModels

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
