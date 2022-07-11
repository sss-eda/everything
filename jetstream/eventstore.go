package jetstream

import (
	"encoding/json"
	"log"

	natsio "github.com/nats-io/nats.go"
)

type AggregateType interface {
	Apply(Event) error
	UnpersistedEvents() []Event
}

type EventStore[Aggregate AggregateType] struct {
	Load func(string) (Aggregate, error)
	Save func(string, Aggregate) error
}

func Load[Aggregate interface{ Apply(Event) error }](
	js natsio.JetStream,
	domain string,
) func(string, string) (Aggregate, error) {
	return func(kind string, id string) (Aggregate, error) {
		var aggregate Aggregate

		js.Subscribe(domain+"."+kind+".events", func(msg *natsio.Msg) {
			event := Event{}
			err := json.Unmarshal(msg.Data, &event)
			if err != nil {
				log.Println(err)
				return
			}
			err = aggregate.Apply(event)
			if err != nil {
				log.Println(err)
				return
			}
		})

		return aggregate, nil
	}
}

func Save[Aggregate interface{ UnpersistedEvents() []Event }](
	js natsio.JetStream,
	domain string,
) func(string, string, Aggregate) error {
	return func(id string, kind string, aggregate Aggregate) error {
		for event := range aggregate.UnpersistedEvents() {
			data, err := json.Marshal(event)
			if err != nil {
				return err
			}
			_, err = js.Publish(domain+"."+kind+".events", data)
			if err != nil {
				return err
			}
		}

		return nil
	}
}
