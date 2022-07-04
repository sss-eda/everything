package es

import "github.com/sss-eda/everything/internal/cqrs"

type Aggregate struct {
	ddd.BaseAggregate
	commandHandlers map[cqrs.CommandKind]cqrs.CommandHandler
}

type AggregateOption func(*Aggregate) error

func WithCommandHandler(
	commandKind cqrs.CommandKind,
	commandHandler cqrs.CommandHandler,
) AggregateOption {
	return func(aggregate *Aggregate) error {
		aggregate.commandHandlers[commandKind] = commandHandler

		return nil
	}
}

func NewAggregate(options ...AggregateOption) (*Aggregate, error) {
	aggregate := Aggregate{}

	var err error

	for _, option := range options {
		err = option(&aggregate)
		if err != nil {
			return nil, err
		}
	}

	return &aggregate, nil
}
