package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/jaquelineabreu/chat-websocket/adapters/websocket"
)

func SetupRouter(wsHandler *websocket.Handler) *gin.Engine {
	r := gin.Default()

	r.StaticFile("/", "./web/index.html")

	r.GET("/ws", func(c *gin.Context) {
		wsHandler.HandleConnections(c.Writer, c.Request)
	})

	return r
}
