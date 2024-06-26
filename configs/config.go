package configs

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type conf struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBName        string `mapstructure:"DB_NAME"`
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
}

func LoadConfig() (*conf, error) {
	var cfg *conf

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(wd)
	viper.SetConfigFile(filepath.Join(wd, "/cmd/pesquisa_de_preco/.env"))
	// viper.SetConfigFile(filepath.Join(wd, ".env"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}
	return cfg, err
}
