package infrastructure

type Fetcher interface {
	FetchRate(baseCurrency string, wantedCurrency string) float64
}
