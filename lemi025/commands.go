package lemi025

import (
	"time"

	"github.com/sss-eda/everything/instrumentation"
)

// type ReadConfigFunc func(context.Context, ReadConfigCommandPayload) error
// type ReadTimeFunc func(context.Context, ReadTimeCommandPayload) error

// type ConfigReadPresenter func(context.Context, chan<- ConfigReadEvent, chan<- error)

// type Subscription interface {
//     Updates() <-chan Event // stream of Events
//     Close() error         // shuts down the stream
// }

// type subscription struct {
// 	closing chan chan error
// 	messages []Message
// }

// func (sub *subscription) loop() {
// 	var err error

// 	for {
// 		select {
// 			case sub.pending
// 			case errc <-sub.closing:
// 				errc <- err
// 				close(sub.updates)
// 				return
// 		}
// 	}
// }

// func (sub *subscription) Updates() <-chan Event {
// 	return nil
// }

// func something(errc chan chan error) {
// 	for {
// 		select {
// 			case err
// 		}
// 	}
// }

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
