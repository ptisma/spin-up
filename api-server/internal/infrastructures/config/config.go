package config

import (
	"flag"
	"os"
)

type Config interface {
	GetApiURL() string
	GetApiPort() string
}

type cfg struct {
	apiURL  string
	apiPort string
}

func (c *cfg) GetApiURL() string {
	return c.apiURL
}
func (c *cfg) GetApiPort() string {
	return c.apiPort
}

func GetConfig() Config {
	conf := &cfg{}
	flag.StringVar(&conf.apiURL, "apiURL", os.Getenv("API_URL"), "API URL")
	flag.StringVar(&conf.apiPort, "apiport", os.Getenv("API_PORT"), "API port")
	return conf

}
