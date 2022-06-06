package es

import "github.com/sss-eda/everything/internal/ddd"

type Store interface {
	Load(ddd.AggregateID) (<-chan ddd.Event, error)
	Save(ddd.AggregateID, ddd.Event) error
}
