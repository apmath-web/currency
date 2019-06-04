package Infrastructure

import (
	"sync"

	"github.com/apmath-web/currency/Domain"
	"github.com/apmath-web/currency/Domain/services"
	"github.com/apmath-web/currency/Infrastructure/applicationModels"
	"github.com/apmath-web/currency/Infrastructure/repositories"
)

type ServiceManagerClass struct {
}

func (sm *ServiceManagerClass) GetRepository() Domain.RepositoryInterface {
	return GenRepository()
}

func (sm *ServiceManagerClass) GetFetcher() Domain.FetcherInterface {
	return applicationModels.GenFetcher()
}

func (sm *ServiceManagerClass) GetExchangerService() Domain.ExchangerInterface {
	service := services.GenExchanger(sm.GetRepository())
	return service
}

func (sm *ServiceManagerClass) GetUpdaterService() Domain.UpdaterInterface {
	service := services.GenUpdater(sm.GetRepository(), sm.GetFetcher())
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

var repo *repositories.Repository

func GenRepository() Domain.RepositoryInterface {
	once.Do(func() {
		repo = &repositories.Repository{Rates: make(map[string]map[string]float64)}
	})
	return repo
}