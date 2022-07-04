package main

import (
	"github.com/sss-eda/everything/jetstream"
	"github.com/sss-eda/everything/lemi025"
	"github.com/sss-eda/everything/serial"

	natsio "github.com/nats-io/nats.go"
	tarm "github.com/tarm/serial"
)

func main() {
	port, _ := tarm.OpenPort(&tarm.Config{Name: "COM1", Baud: 115200})

	adapter, _ := serial.NewLemi025Adapter(port)

	nc, _ := natsio.Connect("wss://nats.sansa.dev")
	js, _ := nc.JetStream()
	store, _ := jetstream.NewEventStore(js)

	lemi025.Connect(
		store,
		cqrs.WithCommandHandler(adapter.ReadConfig),
		cqrs.WithCommandHandler(adapter.ReadTime),
	)

	// aggregate, _ := ddd.NewAggregate(
	// 	lemi025.AggregateKind,

	// 	store,
	// 	es.WithCommandHandler("readConfig", adapter.ReadConfig),
	// 	es.WithCommandHandler("readTime", adapter.ReadTime),
	// 	es.WithCommandHandler("setTime", adapter.SetTime),
	// )

	// nats.Subscribe(nc, "lemi025.001.commands", instrument)
}
