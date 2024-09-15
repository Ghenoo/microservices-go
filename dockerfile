# Etapa de construção
FROM golang:1.22 AS build

# Defina o diretório de trabalho
WORKDIR /app

# Copie go.mod e go.sum e baixe as dependências
COPY go.mod go.sum ./
RUN go mod download

# Copie o código fonte
COPY . .

# Compile o binário
RUN GOARCH=amd64 GOOS=linux go build -o main ./cmd/api

# Etapa final
FROM debian:latest

# Instale as dependências necessárias (certificados CA)
RUN apt-get update && apt-get install -y ca-certificates

# Copie o binário do estágio de construção
COPY --from=build /app/main /app/main

# Defina o diretório de trabalho
WORKDIR /app

# Exponha a porta na qual o app estará rodando
EXPOSE 8080

# Comando para iniciar o aplicativo
CMD ["/app/main"]
