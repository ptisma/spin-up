package serviceContainer

import (
	"api-server/internal/controllers"
	"api-server/internal/infrastructures/application"
	"api-server/internal/repositories"
	"api-server/internal/services"
	"sync"
)

type IServiceContainer interface {
	InjectVMController() controllers.VmController
	InjectDockerController() controllers.DockerController
}

type kernel struct {
	app application.Application
}

func (k *kernel) InjectVMController() controllers.VmController {

	vmService := services.VMService{}

	vmController := controllers.VmController{vmService}

	return vmController
}

func (k *kernel) InjectDockerController() controllers.DockerController {

	dockerBuildJobRepository := &repositories.DockerBuildJobRepository{DB: k.app.GetDB()}

	dockerBuildJobService := &services.DockerBuildJobService{I1: dockerBuildJobRepository}

	dockerController := controllers.DockerController{I1: dockerBuildJobService}

	return dockerController
}

var (
	k             *kernel
	containerOnce sync.Once
)

func GetServiceContainer(app application.Application) IServiceContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{app: app}
		})
	}
	return k
}
