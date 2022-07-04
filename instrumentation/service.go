package instrumentation

import (
	"github.com/sss-eda/everything/internal/cqrs"
	"github.com/sss-eda/everything/internal/ddd"
	"github.com/sss-eda/everything/internal/es"
)

type Service struct {
	eventStore es.EventStore
}

func (service *Service) SendCommand(instrumentID ddd.AggregateID, command cqrs.Command) error {
	instrument, err := service.eventStore.Load(instrumentID)
	if err != nil {
		return err
	}

	err = instrument.SendCommand(command)
	if err != nil {
		return err
	}

	return nil
}

func (service *Service) GetEvents(instrumentID ddd.AggregateID) error {
	instrument, err := service.eventStore.Load(instrumentID)
	if err != nil {
		return err
	}

	instrument.GetEvents()

	return nil
}
