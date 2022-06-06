package lemi025

import "fmt"

type Instrument struct {
	IsRecording bool
	Events      []Event
}

func (instrument *Instrument) ReadConfig() error {
	if instrument.IsRecording {
		return fmt.Errorf("already recording")
	}

}
