package websocket

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/jaquelineabreu/chat-websocket/domain"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Handler struct {
	Broadcast chan domain.Message
	Clients   map[*websocket.Conn]bool
}

func NewHandler(broadcast chan domain.Message) *Handler {
	return &Handler{
		Broadcast: broadcast,
		Clients:   make(map[*websocket.Conn]bool),
	}
}

func (h *Handler) HandleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Erro ao estabelecer conexão WebSocket:", err)
		return
	}
	defer ws.Close()

	h.Clients[ws] = true
	fmt.Println("Cliente conectado.")

	for {
		var msg domain.Message
		if err := ws.ReadJSON(&msg); err != nil {
			fmt.Println("Erro ao ler mensagem:", err)
			delete(h.Clients, ws)
			break
		}

		h.Broadcast <- msg
	}
}
