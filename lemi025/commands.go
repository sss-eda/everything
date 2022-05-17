package lemi025

import (
	"time"

	"github.com/sss-eda/everything/instrument"
)

type Command[C Commands] struct {
	ID      instrument.ID
	Payload C
}

type Commands interface {
	ReadConfigCommand | ReadTimeCommand | SetTimeCommand
}

// ReadConfigCommand TODO
type ReadConfigCommand struct{}

type ReadTimeCommand struct{}

type SetTimeCommand struct {
	Time time.Time
}
