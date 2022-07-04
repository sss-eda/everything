package jetstream

import (
	natsio "github.com/nats-io/nats.go"
	"github.com/sss-eda/everything/internal/ddd"
	"github.com/sss-eda/everything/internal/es"
)

type EventStore struct {
	js natsio.JetStream
}

func NewEventStore(
	js natsio.JetStream,
) (*EventStore, error) {
	return &EventStore{js: js}, nil
}

func (store *EventStore) Save(id ddd.AggregateID, aggregate *es.Aggregate) error {
	return nil
}

func (store *EventStore) Load(id ddd.AggregateID) (*es.Aggregate, error) {
	return nil, nil
}
