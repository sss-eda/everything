package lemi025

import (
	"fmt"
)

type Station struct {
	isRecording bool
	isConnected bool
	adapter     Adapter
}

type Adapter interface {
	ReadConfig(ReadConfigCommand) error
	ReadTime(ReadTimeCommand) error
	// SetTime(SetTimeCommand) error
}

func NewStation() (*Station, error) {
	return &Station{
		isRecording: false,
		isConnected: false,
		adapter:     nil,
	}, nil
}

func (station *Station) Connect(adapter Adapter) error {
	station.adapter = adapter

	station.isConnected = true

	return nil
}

func (station *Station) SendCommand(command Command) error {
	switch c := command.(type) {
	case ReadConfigCommand:
		station.adapter.ReadConfig(c)
	case ReadTimeCommand:
		station.adapter.ReadTime(c)
	default:
		return fmt.Errorf("invalid command type: %v", commandType)
	}
}
