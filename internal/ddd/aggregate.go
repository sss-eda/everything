package ddd

type AggregateID string

type Aggregate interface {
	ID() AggregateID
	Root() Entity
}
