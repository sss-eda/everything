package lemi025

import (
	"time"
)

type BaseEvent struct {
	Timestamp time.Time
	ID        string
}

// type Events interface {
// 	ConfigReadEvent | TimeReadEvent | TimeSetEvent
// }

type ConfigReadEvent struct {
	base    BaseEvent
	payload ConfigReadEventPayload
}

type ConfigReadEventPayload struct {
	StationNumber uint8
}

type TimeReadEventPayload struct {
	Time time.Time
}

type TimeSetEventPayload struct {
	Time time.Time
}
