package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Notification)
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Notification is the notifs channel
type Notification struct {
	Title   string `json:"title"`
	Message string `json:"message"`
}

func handleConnections(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()
	clients[ws] = true
	for {
		// var notif Notification
		// err := ws.ReadJSON(&notif)
		// if err != nil {
		// 	log.Printf("error: %v", err)
		// 	delete(clients, ws)
		// 	break
		// }
		// broadcast <- notif
	}
}

func handleNotifications() {
	for {
		notif := <-broadcast
		for client := range clients {
			err := client.WriteJSON(notif)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func sendNotif(c *gin.Context) {
	body, _ := c.GetRawData()

	var notif Notification
	json.Unmarshal(body, &notif)
	broadcast <- notif
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "NotifyApp API",
			"version": "1.0.0",
		})
	})
	r.POST("/notifs", sendNotif)
	r.GET("/ws", handleConnections)
	go handleNotifications()
	r.Run()
}
