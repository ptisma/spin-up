package config

import (
	"flag"
	"log"
	"os"
)

type Config interface {
	GetApiURL() string
	GetApiPort() string
	GetDBHostURI() string
	GetDBPort() string
}

type cfg struct {
	apiURL    string
	apiPort   string
	dbHostURI string
	dbPort    string
}

func (c *cfg) GetApiURL() string {
	return c.apiURL
}
func (c *cfg) GetApiPort() string {
	return c.apiPort
}
func (c *cfg) GetDBHostURI() string {
	return c.dbHostURI
}
func (c *cfg) GetDBPort() string {
	return c.dbPort
}

func GetConfig() Config {
	conf := &cfg{}
	flag.StringVar(&conf.apiURL, "apiURL", os.Getenv("API_URL"), "API URL")
	flag.StringVar(&conf.apiPort, "apiPort", os.Getenv("API_PORT"), "API port")
	flag.StringVar(&conf.dbHostURI, "dbHostURI", os.Getenv("DB_HOST_URI"), "DB host URI")
	flag.StringVar(&conf.dbPort, "dbPort", os.Getenv("DB_PORT"), "DB port")

	log.Println(conf)
	return conf

}
