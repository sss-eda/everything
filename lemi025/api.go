package lemi025

import (
	"bufio"
	"context"
	"io"
)

type Server interface {
	ReadConfig() error
	ReadTime() error
	SetTime(Time) error
	SetCoefficients1(Mode) error
	ReadCoefficients1() error
	SetCoefficients2(Coefficients2) error
	ReadCoefficients2() error
}

type Client interface {
	io.ReadWriteCloser
}

type EventPublisher interface {
	// OnDatumAcquired(DatumAcquiredEvent)
	OnConfigRead(ConfigReadEvent)
	OnTimeRead(TimeReadEvent)
	OnTimeSet(TimeSetEvent)
}

type EventSubscriber interface {
}

func New(
	ctx context.Context,
	client Client,
	subscriber EventSubscriber,
	publisher EventPublisher,
) (*Service, error) {
	service := Service{
		client:     client,
		subscriber: subscriber,
		publisher:  publisher,
	}

	go service.run(ctx)

	return service, nil
}

type Service struct {
	client     Client
	subscriber EventSubscriber
	publisher  EventPublisher
}

func (service *Service) run(ctx context.Context) error {
	scanner := bufio.NewScanner(service.client)
	scanner.Split(bufio.ScanBytes)

	for scanner.Scan() {
		select {
		case <-ctx.Done():
			break
		}

	}

	return nil
}

func (service *Service) ReadConfig() error {
	_, err := service.client.Write([]byte{0x3D, 0x30})
	if err != nil {
		return err
	}

	return nil
}
