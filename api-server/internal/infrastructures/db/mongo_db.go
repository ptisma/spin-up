package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB interface {
	Close() error
	GetClient() *mongo.Client
}

type database struct {
	Client *mongo.Client
}

func (d *database) GetClient() *mongo.Client {
	return d.Client
}

func (d *database) Close() error {
	var err error
	// TODO
	return err
}

func GetDB(connStr string) (DB, error) {
	var err error
	// Set client options
	clientOptions := options.Client().ApplyURI(connStr)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, err
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB!")

	return &database{
		Client: client,
	}, err
}
