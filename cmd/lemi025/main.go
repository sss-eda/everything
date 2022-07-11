package main

import (
	"net/http"

	"github.com/sss-eda/everything/jetstream"
	"github.com/sss-eda/everything/lemi025/serial"
	tarm "github.com/tarm/serial"
)

func main() {
	// client, _ := rest.NewHTTPClient("https://instrumentation.sansa.dev/")

	port, _ := tarm.OpenPort(&tarm.Config{
		Name: "COM1",
		Baud: 115200,
	})

	nc, _ := natsio.Connect("wss://nats.sansa.dev")
	js, _ := nc.JetStream()

	station, _ := jetstream.Load(js, "instrumentation.lemi025", "station")("39")

	controller := serial.StationController(port, station)
	acquirer := serial.StationAcquirer(port)

	http.HandleFunc("commands", rest.HTTPCommandsHandlerFunc(station))
	http.HandleFunc("events", rest.HTTPEventsHandlerFunc(station))

	http.ListenAndServe(":8080", nil)
}
