package ws

import "github.com/gorilla/websocket"

type Client struct {
	Conn    *websocket.Conn
	Message chan *Message
}

type Message struct {
	Content  string `json:"content"`
	RoomID   string `json:"roomId"`
	Username string `json:"username"`
}
