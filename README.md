# **Chat em Tempo Real com WebSocket e Gin**

Este é um projeto de um **chat em tempo real** desenvolvido em **Golang** utilizando **Gin** e **WebSockets**. O objetivo é criar uma aplicação simples de comunicação em tempo real com uma interface web.

---

## **Funcionalidades**
- Comunicação em tempo real utilizando WebSockets.
- Exibição das mensagens no formato de chat.
- Cada usuário tem uma cor única para identificação visual.
- Interface web simples e funcional.

---


## **Como Executar**

### **1. Requisitos**
Certifique-se de que você possui as seguintes ferramentas instaladas:
- **[Docker](https://www.docker.com/)**

---

### **2. Rodar com Docker**

#### **Construir a Imagem**
No diretório raiz do projeto, execute:
```bash
docker build -t chat-websocket .
```

#### **Iniciar o Container**
Após criar a imagem, execute o seguinte comando para rodar o projeto:


```bash
docker run -p 8081:8081 chat-websocket
```

#### **O servidor será iniciado e estará disponível em:**


```bash
http://localhost:8081/
```






