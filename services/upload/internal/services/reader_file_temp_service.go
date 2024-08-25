package services

import (
	"compras/services/upload/pkg/importer"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

type ReaderFileTempService struct {
	File       *multipart.FileHeader
	OpenedFile multipart.File
}

func (rf *ReaderFileTempService) ReadFileTempService(ctx context.Context) ([][]string, error) {
	ext := filepath.Ext(rf.File.Filename)

	tempFile, err := createdFileTemp(rf.File.Filename, rf.OpenedFile)
	if err != nil {
		return nil, fmt.Errorf("falha ao criar o arquivo temporário: %v", err)
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	data, err := fileScanning(ext, tempFile.Name())
	if err != nil {
		return nil, fmt.Errorf("falha ao escanear o arquivo: %v", err)
	}

	return data, nil
}

func fileScanning(ext string, tempFileName string) ([][]string, error) {
	var data [][]string
	var err error

	switch ext {
	case ".csv":
		data, err = importer.ReadCSVFile(tempFileName)
	case ".xlsx":
		data, err = importer.ReadFileXLSX(tempFileName)
	default:
		err = fmt.Errorf("formato de arquivo não suportado")
	}

	if err != nil {
		return data, fmt.Errorf("falha ao copiar o arquivo: %v", err)
	}

	return data, nil
}

func createdFileTemp(filenametmp string, file multipart.File) (*os.File, error) {
	tempFile, err := os.CreateTemp("/tmp", filenametmp)
	if err != nil {
		return nil, fmt.Errorf("falha ao criar o arquivo temporário: %v", err)
	}

	_, err = io.Copy(tempFile, file)
	if err != nil {
		tempFile.Close()
		os.Remove(tempFile.Name())
		return nil, fmt.Errorf("falha ao copiar o arquivo: %v", err)
	}

	if err := tempFile.Close(); err != nil {
		return nil, fmt.Errorf("falha ao fechar o arquivo temporário: %v", err)
	}

	return tempFile, nil
}
