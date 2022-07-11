package twin

import (
	"fmt"

	"github.com/sss-eda/everything/lemi025"
)

type Station struct {
	isRecording bool
	isConnected bool
	number      lemi025.StationNumber
	changes     []lemi025.Event
	commands    map[CommandID]Command
}

type Command struct {
	sent    bool
	payload any
}

func (station *Station) SendReadConfigCommand(command lemi025.ReadConfigCommand) error {
	if station.isRecording {
		return fmt.Errorf("read-config command not allowed while station is recording data")
	}

	if station.isConnected {
		return fmt.Errorf("cannot send command to disconnected station")
	}

	station.changes = append(station.changes, lemi025.ReadConfigCommandSentEvent{
		Command: command,
	})

	return nil
}

func (station *Station) ApplyEvent(event lemi025.Event) {
	switch event {
	case lemi025.ReadConfigCommandSentEvent:
		station.commands = append(station.commands, event.Command)
	case lemi025.ConfigReadEvent:
		station.
	}
}
