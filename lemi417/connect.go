package lemi417

import "github.com/sss-eda/everything/serial"

type ConnectionKind string

const (
	Serial   ConnectionKind = "serial"
	HttpRest ConnectionKind = "http_rest"
)

// Connection TODO
type Connection interface {
	ReadConfig() error
	ReadTime() error
	Events() <-chan Event
	Close() error
}

// Connect TODO
func Connect(
	kind ConnectionKind,
) (Connection, error) {
	switch kind {
	case Serial:
		return serial.NewLemi025Connection()
	case HttpRest:
		return rest.NewLemi025Connection()
	default:
		return nil, ErrUnknownConnectionKind{}
	}
}
