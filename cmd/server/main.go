package main

import (
	"example.com/websocket-go-2/internal/websocket"
	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	hub := websocket.NewHub()
	websocketHandler := websocket.NewHandler(hub)

	go hub.Run()

	server.POST("/ws/createroom", websocketHandler.CreateRoom)
	server.GET("/ws/joinroom/:roomid", websocketHandler.JoinRoom)
	server.GET("/ws/getrooms", websocketHandler.GetRooms)
	server.GET("/ws/getclients/:roomid", websocketHandler.GetClients)

	server.Run(":3000")
}