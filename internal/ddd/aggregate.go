package ddd

type AggregateID string

type AggregateKind string

type Aggregate interface {
	ID() AggregateID
	Kind() AggregateKind
	Root() Entity
}

func NewAggregate(aggregateKind AggregateKind, aggregateID AggregateID, options ...AggregateOption) (*Aggregate, error) {

}
