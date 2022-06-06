package lemi025

import (
	"fmt"

	"github.com/sss-eda/everything/internal/es"
)

type Service struct {
	events es.Store
}

func NewService(eventStore es.Store) (*Service, error) {
	return &Service{
		events: eventStore,
	}, nil
}

func (service *Service) ReadConfig(
	instrument *Instrument,
) error {
	if instrument.IsRecording() {
		return fmt.Errorf("already recording")
	}

	err := instrument.ReadConfig()
	if err != nil {
		return err
	}

}
