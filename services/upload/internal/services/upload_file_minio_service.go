package services

import (
	"compras/services/upload/config"
	"context"
	"fmt"
	"mime/multipart"

	pkgMinio "compras/services/upload/pkg/minio"

	minio "github.com/minio/minio-go/v7"
)

type UploadFileMinioService struct {
	Ctx         context.Context
	ContentType string
	ObjectName  string
	OpenedFile  multipart.File
	FileSize    int64
}

func (params *UploadFileMinioService) UploadFile() error {

	mn := &pkgMinio.MinioClient{}
	err := mn.Connect(params.Ctx, &pkgMinio.MinioConfig{
		Endpoint:  config.Env.MinioEndpoint,
		AccessKey: config.Env.MinioAccessKey,
		SecretKey: config.Env.MinioSecretKey,
		Bucket:    config.Env.MinioBucket,
	})
	if err != nil {
		return fmt.Errorf("falha ao inicializar o MinIO ou o bucket não existe")
	}
	_, err = mn.SendFile(params.Ctx, config.Env.MinioBucket, params.ObjectName, params.OpenedFile, params.FileSize, minio.PutObjectOptions{
		ContentType: params.ContentType,
	})
	if err != nil {
		return fmt.Errorf("falha ao enviar o arquivo para o MinIO")
	}

	return nil
}
