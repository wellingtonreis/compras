package mongodb

import (
	"fmt"

	"github.com/wellingtonreis/compras/configs"
)

func GetMongoDBURI() string {
	configs, err := configs.LoadConfig()
	if err != nil {
		panic(err)
	}

	// mongodbUri := fmt.Sprintf("%s://%s:%s@%s:%s", configs.DBDriver, configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort)
	mongodbUri := fmt.Sprintf("%s://%s:%s", configs.DBDriver, configs.DBHost, configs.DBPort)
	return mongodbUri
}

func GetMongoSchema() string {
	configs, err := configs.LoadConfig()
	if err != nil {
		panic(err)
	}
	return configs.DBName
}
