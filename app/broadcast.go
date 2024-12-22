package app

import (
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/jaquelineabreu/chat-websocket/domain"
)

func HandleBroadcast(broadcast chan domain.Message, clients map[*websocket.Conn]bool) {
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
