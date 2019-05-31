package domainModels

type CurrencyModel struct {
	name string
}

func (i *CurrencyModel) GetName() string {
	return i.name
}

func GenCurrencyDomainModel(name string) *CurrencyModel {
	return &CurrencyModel{
		name,
	}
}
