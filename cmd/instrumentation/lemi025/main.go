package main

import (
	"log"

	"github.com/sss-eda/everything/pkg/lemi025"
)

func main() {
	log.Fatal(lemi025.Run(
		serial.Connection(),
		websocket.EventBus(),
		websocket.CommandBus(),
	))
}
