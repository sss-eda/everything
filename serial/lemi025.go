package serial

import (
	"io"

	"github.com/sss-eda/everything/lemi025"
)

// func ReadConfig(command lemi025.ReadConfigCommand) error

type Lemi025Adapter struct {
	port io.ReadWriteCloser
}

func NewLemi025Adapter(port io.ReadWriteCloser) (*Lemi025Adapter, error) {
	return &Lemi025Adapter{
		port: port,
	}, nil
}

func (adapter *Lemi025Adapter) ReadConfig(command lemi025.ReadConfigCommand) error {
	_, err := adapter.port.Write([]byte{0x3D, 0x30})
	return err
}

// type Lemi025Adapter struct {
// 	port io.ReadWriteCloser
// }

// // NewLemi025Adapter to a LEMI025
// func NewLemi025Adapter(
// 	port io.ReadWriteCloser,
// ) (*Lemi025Adapter, error) {
// 	adapter := Lemi025Adapter{
// 		port: port,
// 	}

// 	return &adapter, nil
// }

// func (adapter *Lemi025Adapter) ReadConfig(command lemi025.ReadConfigCommand) error {
// 	_, err := adapter.port.Write([]byte{0x3D, 0x30})
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (adapter *Lemi025Adapter) Run() error {
// 	return nil
// }
