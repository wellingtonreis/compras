# Use a imagem base do Go
FROM golang:1.22.1

RUN go install github.com/air-verse/air@latest

ENV PATH="/root/go/bin:${PATH}"

WORKDIR /backend

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go mod tidy

RUN go build -o bin/compras cmd/pesquisa_de_preco/main.go

CMD ["air"]

EXPOSE 3000
