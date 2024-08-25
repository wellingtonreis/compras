package minio

import (
	"context"
	"log"
	"mime/multipart"

	minio "github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioClient struct {
	Client *minio.Client
}

type MinioConfig struct {
	Endpoint  string
	AccessKey string
	SecretKey string
	Bucket    string
}

func (mn *MinioClient) Connect(ctx context.Context, config *MinioConfig) error {

	log.Println("Iniciando conexão com MinIO...")

	var err error
	mn.Client, err = minio.New(config.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKey, config.SecretKey, ""),
		Secure: false,
	})
	if err != nil {
		log.Printf("Falha ao inicializar o MinIO: %v", err)
		return err
	}

	_, err = mn.Client.BucketExists(ctx, config.Bucket)
	if err != nil {
		log.Printf("A conexão foi realizada mas o bucket não existe: %v", err)
		return err
	}

	return nil
}

func (mn *MinioClient) SendFile(ctx context.Context, bucketName, objectName string, reader multipart.File, size int64, opts minio.PutObjectOptions) (string, error) {

	log.Println("Realizando upload do arquivo...")

	_, err := mn.Client.PutObject(ctx, bucketName, objectName, reader, size, opts)
	if err != nil {
		log.Printf("Falha ao enviar o arquivo para o MinIO: %v", err)
		return "", err
	}

	log.Println("Upload do arquivo realizado com sucesso!")

	return objectName, nil
}
