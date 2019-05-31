package Infrastructure

import (
	"github.com/apmath-web/currency/Domain"
	"github.com/apmath-web/currency/Infrastructure/applicationModels"
)

type ServiceManager struct {
}

func (sm *ServiceManager) GetFetcherService() Domain.Fetcher {
	return applicationModels.GenFetcherApplicationModel()
}
