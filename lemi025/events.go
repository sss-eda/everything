package lemi025

import (
	"time"

	"github.com/sss-eda/everything/instrumentation"
)

type Event[E Events] struct {
	InstrumentID instrumentation.InstrumentID
	Payload      E
}

type Events interface {
	ConfigReadEvent | TimeReadEvent | TimeSetEvent
}

type ConfigReadEvent struct {
	StationNumber uint8
}

type TimeReadEvent struct {
	Time time.Time
}

type TimeSetEvent struct {
	Time time.Time
}

type ConfigRead struct{}
