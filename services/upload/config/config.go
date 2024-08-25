package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type conf struct {
	DBDriver         string `mapstructure:"DB_DRIVER"`
	DBHost           string `mapstructure:"DB_HOST"`
	DBPort           string `mapstructure:"DB_PORT"`
	DBUser           string `mapstructure:"DB_USER"`
	DBPassword       string `mapstructure:"DB_PASSWORD"`
	DBName           string `mapstructure:"DB_NAME"`
	WebServerPort    string `mapstructure:"WEB_SERVER_PORT"`
	GovbrUrl         string `mapstructure:"URL_SRC_GOVBR"`
	MinioEndpoint    string `mapstructure:"MINIO_ENDPOINT"`
	MinioAccessKey   string `mapstructure:"MINIO_ACCESS_KEY"`
	MinioSecretKey   string `mapstructure:"MINIO_SECRET_KEY"`
	MinioBucket      string `mapstructure:"MINIO_BUCKET"`
	RabbitMQUrl      string `mapstructure:"RABBITMQ_URL"`
	RabbitMQQueue    string `mapstructure:"RABBITMQ_QUEUE"`
	RabbitMQExchange string `mapstructure:"RABBITMQ_EXCHANGE"`
}

var Env *conf

func LoadConfig() (*conf, error) {
	path_file_env := os.Getenv("PATH_ROOT")

	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path_file_env)
	viper.SetConfigFile(filepath.Join(path_file_env, ".env"))
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}
	err = viper.Unmarshal(&Env)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return Env, err
}
