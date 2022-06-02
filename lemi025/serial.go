package lemi025

import "io"

// Connection TODO
type Connection struct {
	clients []Client
}

// Connect TODO
func Connect(options ...ConnectOption) (Connection, error) {

}

type ConnectOption func(Connection) error

func WithSerial(port io.ReadWriteCloser) ConnectOption {
	return func(conn Connection) error {

	}
}
