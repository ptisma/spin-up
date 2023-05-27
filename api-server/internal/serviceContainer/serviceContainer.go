package serviceContainer

import (
	"api-server/internal/controllers"
	"api-server/internal/infrastructures/application"
	"api-server/internal/services"
	"sync"
)

type IServiceContainer interface {
	InjectVMController() controllers.VmController
}

type kernel struct{}

func (k *kernel) InjectVMController() controllers.VmController {

	vmService := services.VMService{}

	vmController := controllers.VmController{vmService}

	return vmController
}

var (
	k             *kernel
	containerOnce sync.Once
)

func GetServiceContainer(app application.Application) IServiceContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
		})
	}
	return k
}
