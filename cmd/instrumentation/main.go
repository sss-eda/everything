package main

import "log"

func main() {
	eventStore := nats.NewEventStore()

	service := instrumentation.NewService(eventStore)

	apiServer := rest.NewInstrumentationAPI(service)

	log.Fatal(apiServer.Run())
}
