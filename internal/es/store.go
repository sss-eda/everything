package es

import "github.com/sss-eda/everything/internal/ddd"

type EventStore interface {
	Save(ddd.AggregateType, ddd.AggregateID, *ddd.Aggregate) error
	Load(ddd.AggregateType, ddd.AggregateID) (*ddd.Aggregate, error)
}
