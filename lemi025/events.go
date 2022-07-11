package lemi025

import "time"

type ReadConfigCommandSentEvent struct {
	Command ReadConfigCommand
}

type ConfigReadEvent struct {
	StationNumber StationNumber
}

type TimeReadEvent struct {
	Time time.Time
}
