package application

import (
	"api-server/internal/infrastructures/config"
	"api-server/internal/infrastructures/db"
	"fmt"
)

type Application interface {
	GetConfig() config.Config
	GetDB() db.DB
}

type application struct {
	Config config.Config
	DB     db.DB
}

func (a *application) GetConfig() config.Config {
	return a.Config
}

func (a *application) GetDB() db.DB {
	return a.DB
}

func GetApplication() (Application, error) {
	var err error
	config := config.GetConfig()
	if err != nil {
		return nil, err
	}

	dbConnStr := fmt.Sprintf("mongodb://%s:%s", config.GetDBHostURI(), config.GetDBPort())

	db, err := db.GetDB(dbConnStr)
	if err != nil {
		return nil, err
	}

	return &application{
		Config: config,
		DB:     db,
	}, err
}
