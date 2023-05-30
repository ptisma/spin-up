package router

import (
	"api-server/internal/infrastructures/application"
	"api-server/internal/serviceContainer"
	"sync"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type IMuxRouter interface {
	InitRouter() *mux.Router
}

type router struct {
	app application.Application
}

func (router *router) InitRouter() *mux.Router {

	serviceContainer := serviceContainer.GetServiceContainer(router.app)

	vmController := serviceContainer.InjectVMController()
	dockerController := serviceContainer.InjectDockerController()

	mux := mux.NewRouter()

	// mux.Use(middlewares.Auth0Middleware, middlewares.PrometheusMiddleware)

	mux.Path("/prometheus").Handler(promhttp.Handler())

	mux.HandleFunc("/health", vmController.GetHealth).Methods("GET")
	mux.HandleFunc("/docker-img-build", dockerController.CreateDockerBuildJob).Methods("POST")

	return mux
}

var (
	m          *router
	routerOnce sync.Once
)

func GetMuxRouter(app application.Application) IMuxRouter {
	if m == nil {
		routerOnce.Do(func() {
			m = &router{app: app}
		})
	}
	return m
}
