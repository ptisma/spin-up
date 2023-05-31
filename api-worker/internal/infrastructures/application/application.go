package application

import (
	"api-worker/internal/infrastructures/broker"
	"api-worker/internal/infrastructures/config"
)

type Application interface {
	GetConfig() config.Config
	GetBroker() broker.Broker
}

type application struct {
	Config config.Config
	Broker broker.Broker
}

func (a *application) GetConfig() config.Config {
	return a.Config
}

func (a *application) GetBRoker() broker.Broker {
	return a.Broker
}

func GetApplication() (Application, error) {
	var err error
	config := config.GetConfig()
	if err != nil {
		return nil, err
	}
	broker := broker.GetBroker()
	if err != nil {
		return nil, err
	}
	return &application{
		Config: config,
		Broker: broker,
	}, err
}
