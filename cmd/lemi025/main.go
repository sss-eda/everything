package main

import "log"

func main() {
	client, err := websocket.NewInstrumentClient("ws://api.sansa.dev/instrument")
	if err != nil {
		log.Fatal(err)
	}

	instrument, err := serial.NewLemi025Adapter("/dev/ttyACM0")

	client.Subscribe(lemi025.ReadConfigCommand, lemi025.ReadConfig)
	client.Subscribe(lemi025.ReadTimeCommand, lemi025.ReadTime)
	client.Subscribe(lemi025.SetTimeCommand, lemi025.SetTime)

	instrument.Subscribe(lemi025.ConfigReadEvent, client.Publish)
	instrument.Subscribe(lemi025.TimeReadEvent, client.Publish)
}
