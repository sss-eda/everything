package main

import (
	"net/http"

	natsio "github.com/nats-io/nats.go"
	"github.com/sss-eda/everything/instrumentation"
	"github.com/sss-eda/everything/jetstream"
)

func main() {
	nc, _ := natsio.Connect("wss://instrumentation.sansa.dev/")
	js, _ := nc.JetStream()

	instruments := jetstream.EventStore[instrumentation.Instrument]{
		jetstream.Load(js, "instrumentation.instruments"),
		jetstream.Save(js, "instrumentation.instruments"),
	}

	// The use case: Returns a function, requires a repository
	instrumentation.RegisterInstrument(instruments)

	http.HandleFunc("/instruments", rest.Instruments(
		registration.RegisterInstrumentFunc(jetstream.Load(js), jetstream.Save(js)),
		listing.ListInstrumentsFunc(jetstream.Load(js)),
	))

	http.HandleFunc("/instruments/:id/commands", rest.CommandsHTTPHandleFunc(
		listing.ListInstrumentCommandsFunc()
	))

	// what to use as eventstore...
	// Now what is the bounded context
	// Aggregates
}
