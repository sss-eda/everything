package websocket

import (
	gorillaws "github.com/gorilla/websocket"
)

func Subscribe(
	done chan any,
	conn *gorillaws.Conn,
	streams ...string,
) <-chan []byte {
	output := make(chan []byte)

	go func() {
		defer close(output)

		for {
			select {
			case <-done:
				return
			case message := <-input:
				for _, stream := range streams {
					if err := conn.WriteMessage(gorillaws.TextMessage, message); err != nil {
						return
					}
				}
			}
		}
	}
}
