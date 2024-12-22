# Imagem base do Golang
FROM golang:1.22

# Diretório de trabalho dentro do container
WORKDIR /app

# Copia os arquivos go.mod e go.sum para o container
COPY go.mod go.sum ./

# Instala as dependências
RUN go mod download

# Copia o restante do código fonte para o container
COPY . .

# Compila o binário
RUN go build -o main cmd/main.go

# Porta exposta pelo container
EXPOSE 8081

# Comando para rodar o servidor
CMD ["./main"]
