package websocket

import (
	"github.com/gorilla/websocket"
)

type MessageBus struct {
	ws *websocket.Conn
}

func (bus *MessageBus) Subscribe() {
	ws.
}

func (bus *MessageBus) Publish() {
	
}

// // Publish TODO
// func (connection *InstrumentConnection) Publish(message []byte) error {
// 	return connection.ws.WriteMessage(websocket.TextMessage, message)
// }

// // Subscribe TODO
// func (connection *InstrumentConnection) Subscribe(symbol string) error {
// 	connection.ws.

// }
