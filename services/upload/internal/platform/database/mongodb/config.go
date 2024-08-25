package mongodb

import (
	"compras/services/upload/config"
	"fmt"
)

func GetMongoDBURI() string {
	config, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	mongodbUri := fmt.Sprintf("%s://%s:%s", config.DBDriver, config.DBHost, config.DBPort)
	return mongodbUri
}

func GetMongoSchema() string {
	config, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	return config.DBName
}
