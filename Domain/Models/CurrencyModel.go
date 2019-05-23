package domainModels

type CurrencyModel struct {
	name string
}

func (i CurrencyModel) GetName() string {
	return i.name
}
