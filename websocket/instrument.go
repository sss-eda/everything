package websocket

import (
	"github.com/gorilla/websocket"
)

// InstrumentConnection TODO
type InstrumentConnection struct {
	ws *websocket.Conn
}

// Publish TODO
func (connection *InstrumentConnection) Publish(message []byte) error {
	return connection.ws.WriteMessage(websocket.TextMessage, message)
}

// Subscribe TODO
func (connection *InstrumentConnection) Subscribe(symbol string) error {
	connection.ws.

}

// NewInstrumentConnection TODO
func NewInstrumentConnection(url string) (*InstrumentConnection, error) {
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return nil, err
	}

	return &InstrumentConnection{ws: conn}, nil
}
