package instrument

import (
	"github.com/sss-eda/everything/lemi025"
	"github.com/sss-eda/everything/lemi417"
)

type InstrumentKind string

const (
	Lemi025 InstrumentKind = "LEMI025"
	Lemi417 InstrumentKind = "LEMI417"
)

type Connection interface {
	Subscribe(chan<- Message)
	Publish(Message)
	Close() error
}

// Connect TODO
func Connect(
	kind InstrumentKind,
) (Connection, error) {
	switch kind {
	case Lemi025:
		lemi025.Connect("serial")
	case Lemi417:
		lemi417.Connect("serial")
	}
}
