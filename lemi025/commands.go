package lemi025

import (
	"time"

	"github.com/sss-eda/everything/instrumentation"
)

type Command[C Commands] struct {
	InstrumentID instrumentation.InstrumentID
	Payload      C
}

type Commands interface {
	ReadConfigCommand | ReadTimeCommand | SetTimeCommand | SetCoefficients1Command |
		ReadCoefficients1Command | SetCoefficients2Command | ReadCoefficients2Command |
		ReadGPSDataCommand | StopSystemCommand | StartSystemCommand | CheckFLASHCommand |
		SetDACxCommand | SetDACyCommand | SetDACzCommand
}

type ReadConfigCommand struct{}

func (command ReadConfigCommand) Publish(
	id InstrumentID,
) error {
	return nil
}

type ReadTimeCommand struct{}

type SetTimeCommand struct {
	Time time.Time
}

type SetCoefficients1Command struct {
	Mode Mode
}

type ReadCoefficients1Command struct{}

type SetCoefficients2Command struct {
	Coefficients [20]float32
}

type ReadCoefficients2Command struct{}

type ReadGPSDataCommand struct{}

type StopSystemCommand struct{}

type StartSystemCommand struct{}

type CheckFLASHCommand struct{}

type SetDACxCommand struct {
	DACx int16
}

type SetDACyCommand struct {
	DACy int16
}

type SetDACzCommand struct {
	DACz int16
}
