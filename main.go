package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var broadcast = make(chan Message)
var clients = make(map[*websocket.Conn]bool)

func handleConnections(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Erro ao estabelecer conexão WebSocket:", err)
		return
	}
	defer ws.Close()

	clients[ws] = true

	for {
		var msg Message
		if err := ws.ReadJSON(&msg); err != nil {
			fmt.Println("Erro ao ler mensagem:", err)
			delete(clients, ws)
			break
		}

		broadcast <- msg
	}
}

func handleMessage() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				fmt.Println("Erro ao enviar mensagem:", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func main() {
	r := gin.Default()

	r.StaticFile("/", "./index.html")

	r.GET("/ws", handleConnections)

	go handleMessage()

	// Iniciar o servidor
	fmt.Println("Servidor rodando em http://localhost:8081")
	if err := r.Run(":8081"); err != nil {
		panic(err)
	}
}
