package websocket

import (
	"encoding/json"

	"github.com/gorilla/websocket"
)

// Publish TODO
func Publish[Payload any](conn *websocket.Conn) func(Payload) error {
	return func(payload Payload) error {
		writeCloser, err := conn.NextWriter(websocket.BinaryMessage)
		if err != nil {
			return err
		}
		defer writeCloser.Close()

		message, err := json.Marshal(payload)
		if err != nil {
			return err
		}

		_, err = writeCloser.Write(message)

		return err
	}
}
