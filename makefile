# Nome do binário gerado
BINARY_NAME=chat-websocket

# Caminho do arquivo principal do projeto
MAIN_FILE=cmd/main.go

# Variáveis auxiliares
BUILD_DIR=build

# Comandos
.PHONY: all run build clean tidy

# Rodar o projeto
run:
	@echo "Iniciando o servidor..."
	go run $(MAIN_FILE)

# Instalar dependências
tidy:
	@echo "Instalando dependências..."
	go mod tidy

# Compilar o projeto
build: tidy
	@echo "Compilando o binário..."
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_FILE)

# Limpar arquivos gerados
clean:
	@echo "Limpando arquivos gerados..."
	rm -rf $(BUILD_DIR)

# Rodar o binário compilado
start: build
	@echo "Iniciando o projeto compilado..."
	./$(BUILD_DIR)/$(BINARY_NAME)
