package lemi025

import (
	"testing"

	"github.com/sss-eda/everything/websocket"
)

func TestRun(t *testing.T) {
	done := make(chan any)
	websocket.Publish(
		done,
		Run(
			done,
			websocket.Subscribe(
				done,
				"everything.lemi025.1",
			),
		),
	)
}
