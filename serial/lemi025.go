package serial

import (
	"context"

	"github.com/sss-eda/everything/lemi025"
	tarm "github.com/tarm/serial"
)

// Lemi025Adapter TODO
type Lemi025Adapter struct {
	port *tarm.Port
}

// NewLemi025Adapter TODO
func NewLemi025Adapter(name string, baud int) (*Lemi025Adapter, error) {
	port, err := tarm.OpenPort(&tarm.Config{
		Name: name,
		Baud: baud,
	})
	if err != nil {
		return nil, err
	}

	adapter := &Lemi025Adapter{port: port}

	go adapter.loop()

	return adapter, nil
}

func (adapter *Lemi025Adapter) Run(ctx context.Context) error {
	for {
		// TODO
	}
}

// ReadConfig TODO
func (adapter *Lemi025Adapter) ReadConfig(command lemi025.ReadConfigCommand) error {
	_, err := adapter.port.Write([]byte{0x3D, 0x30})
	if err != nil {
		return err
	}

	return nil
}
