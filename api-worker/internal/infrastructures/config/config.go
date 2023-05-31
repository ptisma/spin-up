package config

import (
	"flag"
	"os"
)

type Config interface {
	GetBrokerURI() string
	GetExchangeName() string
	GetRoutingKey() string
	GetQueueName() string
}

type cfg struct {
	BrokerURI    string
	ExchangeName string
	RoutingKey   string
	QueueName    string
}

func (c *cfg) GetBrokerURI() string {
	return c.BrokerURI
}
func (c *cfg) GetExchangeName() string {
	return c.ExchangeName
}
func (c *cfg) GetRoutingKey() string {
	return c.RoutingKey
}
func (c *cfg) GetQueueName() string {
	return c.QueueName
}

func GetConfig() Config {
	conf := &cfg{}
	flag.StringVar(&conf.BrokerURI, "BrokerURI", os.Getenv("BROKER_URI"), "Broker URI")
	flag.StringVar(&conf.ExchangeName, "ExchangeName", os.Getenv("EXCHANGE_NAME"), "Exchange Name")
	flag.StringVar(&conf.RoutingKey, "RoutingKey", os.Getenv("ROUTING_KEY"), "Routing Key")
	flag.StringVar(&conf.QueueName, "QueueName", os.Getenv("QUEUE_NAME"), "Queue Name")
	return conf

}
