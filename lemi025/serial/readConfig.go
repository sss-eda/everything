package serial

import (
	"bufio"
	"io"

	"github.com/sss-eda/everything/lemi025"
)

type Station struct {
	isRecording bool
	number      lemi025.StationNumber
	history     []Event
}

func NewStation(reader io.Reader) *Station {
	station := Station{
		isRecording: false,
	}

	go func() {
		scanner := bufio.NewScanner(reader)
		station.evaluate()
	}()

	return &station
}

func (station *Station) evaluate() {

	for {

	}
}

func (station *Station) Number() lemi025.StationNumber {
	return station.number
}

func (station *Station) ReadConfig(w io.Writer) func(lemi025.ReadConfigCommand) error {
	return func(command lemi025.ReadConfigCommand) error {
		if station.isRecording {
			return ErrStationRecording{}
		}

		_, err := w.Write([]byte{0x3D, 0x30})

		return err
	}
}

func (station *Station) ReadTime(w io.Writer) func(lemi025.ReadTimeCommand) error {
	return func(command lemi025.ReadTimeCommand) error {
		if station.isRecording {
			return ErrStationRecording{}
		}

		_, err := w.Write([]byte{0x3D, 0x31})

		return err
	}
}
