package main

import (
	"github.com/jaquelineabreu/chat-websocket/adapters/gin"
	"github.com/jaquelineabreu/chat-websocket/adapters/websocket"
	"github.com/jaquelineabreu/chat-websocket/app"
	"github.com/jaquelineabreu/chat-websocket/domain"
)

func main() {
	broadcast := make(chan domain.Message)

	wsHandler := websocket.NewHandler(broadcast)

	go app.HandleBroadcast(broadcast, wsHandler.Clients)

	r := gin.SetupRouter(wsHandler)

	if err := r.Run(":8081"); err != nil {
		panic(err)
	}
}
