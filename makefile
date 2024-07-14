# Nome do binário
BINARY_NAME=compras

# Diretório do projeto
PROJECT_DIR=/home/wgreis/go/src/compras

# Diretório de origem e destino do build
SRC_DIR=$(PROJECT_DIR)
BUILD_DIR=$(PROJECT_DIR)/bin
FILE_INIT=$(PROJECT_DIR)/cmd/pesquisa_de_preco/main.go

# Configurações de compilação
GOOS=linux
GOARCH=amd64

# Comandos padrões
.PHONY: all build run test clean deps tidy vendor fmt vet install

all: build

build:
	@echo "==> Building the project..."
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(BUILD_DIR)/$(BINARY_NAME) $(SRC_DIR)

run:
	@echo "==> Running the project..."
	go run $(FILE_INIT)

test:
	@echo "==> Running tests..."
	cd $(SRC_DIR) && go test ./...

clean:
	@echo "==> Cleaning up..."
	rm -rf $(BUILD_DIR)/*

deps:
	@echo "==> Downloading dependencies..."
	cd $(SRC_DIR) && go get ./...

tidy:
	@echo "==> Tidying up modules..."
	cd $(SRC_DIR) && go mod tidy

vendor:
	@echo "==> Vendoring dependencies..."
	cd $(SRC_DIR) && go mod vendor

fmt:
	@echo "==> Formatting code..."
	cd $(SRC_DIR) && go fmt ./...

vet:
	@echo "==> Vetting code..."
	cd $(SRC_DIR) && go vet ./...

install:
	@echo "==> Installing the package..."
	cd $(SRC_DIR) && go install ./...

# Comando para ver as variáveis
vars:
	@echo "==> Displaying variables..."
	@echo "BINARY_NAME = $(BINARY_NAME)"
	@echo "PROJECT_DIR = $(PROJECT_DIR)"
	@echo "SRC_DIR = $(SRC_DIR)"
	@echo "BUILD_DIR = $(BUILD_DIR)"
	@echo "GOOS = $(GOOS)"
	@echo "GOARCH = $(GOARCH)"
