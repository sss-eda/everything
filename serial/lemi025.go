package serial

import (
	"bufio"
	"context"
	"io"

	"github.com/sss-eda/everything/lemi025"
)

// Lemi025Adapter TODO
type Lemi025Adapter struct{}

// Control TODO
func (adapter *Lemi025Adapter) Control(
	events chan<- lemi025.Event,
) func(context.Context, io.Reader) error {
	return func(ctx context.Context, reader io.Reader) error {
		scanner := bufio.NewScanner(reader)
		scanner.Split(bufio.ScanBytes)

		for scanner.Scan() {
			select {
			case <-ctx.Done():
				break
			}
		}

		return ctx.Err()
	}
}

// ReadConfig TODO
func (adapter *Lemi025Adapter) ReadConfig(
	writer io.Writer,
) func(context.Context) error {
	return func(ctx context.Context) error {
		_, err := writer.Write([]byte{0x3D, 0x30})
		if err != nil {
			return err
		}

		return nil
	}
}

func (adapter *Lemi025Adapter) ReadTime() error {
	_, err := presenter.port.Write([]byte{0x3D, 0x31})
	if err != nil {
		return err
	}

	return nil
}

func (adapter *Lemi025Adapter) SetTime(time lemi025.Time) error {
	return nil
}

instrument := serial.Lemi025Adapter{}

instrument.Control()(context.Background(), port)

instrument.ReadConfig(port)