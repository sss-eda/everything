package instrumentation

// Message TODO
type Message interface {
	Kind() MessageKind
	Payload() []byte
}

type MessageKind string
