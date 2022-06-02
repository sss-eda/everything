package es

type Event interface {
	AggregateKind() string
	AggregateID() string
	Version() uint
	Sequence() uint
	Payload() []byte
}
