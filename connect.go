package main

// ブラウザから送信されたデータを格納
type WsPayload struct {
	Action string
	ID     int
}
