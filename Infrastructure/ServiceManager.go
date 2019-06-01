package Infrastructure

import (
	"sync"

	"github.com/apmath-web/currency/Infrastructure/applicationModels"

	"github.com/apmath-web/currency/Domain"
	"github.com/apmath-web/currency/Domain/services"
	"github.com/apmath-web/currency/Infrastructure/repositories"
)

type ServiceManagerClass struct {
}

func (sm *ServiceManagerClass) GetRepository() Domain.ChangeTableRepositoryInterface {
	return repositories.GenRepository()
}

func (sm *ServiceManagerClass) GetFetcherService() Domain.Fetcher {
	service := applicationModels.GenFetcherApplicationModel()
	return service
}

func (sm *ServiceManagerClass) GetUpdateService() Domain.UpdaterTable {
	service := services.GenUpdaterTableDomainModel(sm.GetFetcherService())
	return service
}

var serviceManager *ServiceManagerClass
var once sync.Once

func GetServiceManager() *ServiceManagerClass {
	once.Do(func() {
		serviceManager = &ServiceManagerClass{}
	})
	return serviceManager
}
