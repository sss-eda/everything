package serial

import (
	"bufio"
	"io"
	"log"

	"github.com/sss-eda/everything/lemi025"
)

type Lemi025Connection struct {
	port    io.ReadWriteCloser
	closing chan chan error
	events  chan lemi025.Event
}

// NewLemi025Connection to a LEMI025
func NewLemi025Connection(port io.ReadWriteCloser) (*Lemi025Connection, error) {
	sub := Lemi025Connection{
		port:    port,
		closing: make(chan chan error),
	}

	go sub.loop()

	return &sub, nil
}

func (conn *Lemi025Connection) loop() {
	var err error
	var buffer []byte
	var state string = "idle"

	scanner := bufio.NewScanner(conn.port)
	scanner.Split(bufio.ScanBytes)

	for {
		select {
		case errc := <-conn.closing:
			errc <- err
			close(conn.closing)
			return
		case event <- events:
			conn.events <- event
		default:
			if scanner.Scan() {
				b := scanner.Bytes()[0]
				buffer = append(buffer, b)

				for {
					switch state {
					case "idle":
						switch buffer[0] {
						case 0x3F:
							state = "response"
						case 0x4C:
							state = "data"
						default:
							buffer = buffer[1:]
							log.Println("drop byte: %x", b)
						}
					case "response":
						switch buffer[0] {
						case 0x30:
							state = "readConfig"
						case 0x31:
							state = "readTime"
						case 0x32:
							state = "setTime"
						default:
							state = "idle"
						}
					case "data":
					}
				}
			}
		}
	}
}

// Events TODO
func (conn *Lemi025Connection) Events() <-chan lemi025.Event {
	return conn.events
}

// Close the subscription and return the error (if any).
func (conn *Lemi025Connection) Close() error {
	errc := make(chan error)
	conn.closing <- errc
	return <-errc
}
