package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handleConnections(h *hub, c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	client := &client{hub: h, conn: conn, send: make(chan []byte, 256)}
	client.hub.register <- client

	go client.write()
	go client.read()
}

func main() {
	r := gin.Default()
	hub := newHub()
	go hub.run()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "NotifyApp API",
			"version": "1.0.0",
		})
	})
	r.GET("/ws", func(c *gin.Context) {
		handleConnections(hub, c)
	})
	r.Run()
}
