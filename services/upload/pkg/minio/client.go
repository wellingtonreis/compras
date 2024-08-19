package minio

import (
	"context"
	"time"

	"compras/services/upload/config"

	minio "github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var MinioClient *minio.Client

func InitMinio() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	MinioClient, err = minio.New(config.MinioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.MinioAccessKey, config.MinioSecretKey, ""),
		Secure: false,
	})
	if err != nil {
		return err
	}

	_, err = MinioClient.BucketExists(ctx, config.MinioBucket)
	if err != nil {
		return err
	}

	return nil
}
