package application

import "api-server/internal/infrastructures/config"

type Application interface {
	GetConfig() config.Config
}

type application struct {
	Config config.Config
}

func (a *application) GetConfig() config.Config {
	return a.Config
}

func GetApplication() (Application, error) {
	var err error
	config := config.GetConfig()
	if err != nil {
		return nil, err
	}
	return &application{
		Config: config,
	}, err
}
