package main

import "github.com/gorilla/websocket"

// WebSocketsコネクション情報を格納
type WebSocketConnection struct {
	*websocket.Conn
}

// ブラウザから送信されたデータを格納
type WsPayload struct {
	Action string
	ID     int
}
