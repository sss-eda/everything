package websocket

import "github.com/gorilla/websocket"

func Subscribe(conn *websocket.Conn) func(string) error {
	return func(id InstrumentID) error {
		return conn.WriteJSON(id)
	}
}
