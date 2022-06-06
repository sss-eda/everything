package lemi025

import (
	"testing"

	gorillaws "github.com/gorilla/websocket"
)

func TestRun(t *testing.T) {
	done := make(chan any)

	wsConn, _, err := gorillaws.DefaultDialer.Dial("ws://sansa.dev/ws", nil)
	if err != nil {
		t.Fatal(err)
	}

	websocket.Publish(
		done,
		wsConn,
		Run(
			done,
			websocket.Subscribe(
				done,
				wsConn,
				"everything.lemi025.1",
			),
		),
		"everything.lemi025.1",
	)
}
