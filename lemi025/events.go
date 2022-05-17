package lemi025

import (
	"time"

	"github.com/sss-eda/everything/instrument"
)

// Event TODO
type Event[E Events] struct {
	ID      instrument.ID
	Payload E
}

type Events interface {
	ConfigReadEvent | TimeReadEvent | TimeSetEvent
}

// ConfigReadEvent TODO
type ConfigReadEvent struct {
	StationNumber uint8
}

type TimeReadEvent struct {
	Time time.Time
}

type TimeSetEvent struct {
	Time time.Time
}
