package listing

import "log"

type QueryService struct {
	Instruments InstrumentRepository
}

func (service *Service) GetAllInstruments() ([]Instrument, error) {
	return service.Instruments.GetAllInstruments()
}

func (service *Service) GetInstrumentByID(id string) (Instrument, error) {
	return service.Instruments.GetInstrumentByID(id)
}

func (service *Service) OnEvent(event Event) {
	switch event.Kind {
	case "instrument-registered":
		service.Instruments.AddInstrument(event.Payload)
	case "command-sent-to-instrument":
		instrument, err := service.Instruments.GetInstrumentByID(event.AggregateID)
		if err != nil {
			log.Println(err)
			return
		}
		switch event.Kind
	}
}

// This is a command service
// func (service *Service) RegisterInstrument(instrumentID string) error {
// 	err := service.Instruments.CreateInstrument(instrumentID)
// 	if err != nil {
// 		return err
// 	}

// 	instrument, err := service.Instruments.GetInstrumentByID(instrumentID)
// 	if err != nil {
// 		return err
// 	}

// 	err = instrument.Register()

// 	return err
// }

// THis is also a command service
// func (service *Service) SendCommandToInstrument(instrumentID string, command Command) error {
// 	instrument, err := service.Instruments.GetInstrumentByID(instrumentID)
// 	if err != nil {
// 		return err
// 	}

// 	err = instrument.SendCommand(command)

// 	return nil
// }
