package viewModels

type CurrencyViewModel struct {
	Amount          int    `json:"amount"`
	CurrentCurrency string `json:"currentCurrency"`
	WantedCurrency  string `json:"wantedCurrency"`
	validation      interface{}
}
