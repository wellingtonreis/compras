package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

var (
	MinioEndpoint  string
	MinioAccessKey string
	MinioSecretKey string
	MinioBucket    string
)

func Minio() {
	path_file_env := os.Getenv("PATH_ROOT")

	err := godotenv.Load(filepath.Join(path_file_env, ".env"))
	if err != nil {
		log.Fatal("Não foi possível carregar o arquivo .env de MinIO: %v ", err)
	}

	MinioEndpoint = os.Getenv("MINIO_ENDPOINT")
	MinioAccessKey = os.Getenv("MINIO_ACCESS_KEY")
	MinioSecretKey = os.Getenv("MINIO_SECRET_KEY")
	MinioBucket = os.Getenv("MINIO_BUCKET")

	if MinioBucket == "" {
		MinioBucket = "my-bucket"
	}
}
