package main

import (
	"log"

	"github.com/sss-eda/everything/lemi025"
)

func main() {
	client, err := websocket.NewInstrumentClient("ws://api.sansa.dev/instrument")
	if err != nil {
		log.Fatal(err)
	}

	instrument, err := serial.NewLemi025Adapter("/dev/ttyACM0")

	commands := client.Subscribe("commands")
	events := instrument.Subscribe("events")

	client.Subscribe(lemi025.ReadConfigCommand, instrument.ReadConfig)
	client.Subscribe(lemi025.ReadTimeCommand, instrument.ReadTime)
	client.Subscribe(lemi025.SetTimeCommand, instrument.SetTime)

	instrument.Subscribe(lemi025.ConfigReadEvent, client.Publish)
	instrument.Subscribe(lemi025.TimeReadEvent, client.Publish)
}
