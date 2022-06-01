package lemi025

import (
	"testing"

	tarm "github.com/tarm/serial"
)

func TestConnect(t *testing.T) {
	port, err := tarm.OpenPort(&tarm.Config{
		Name: "COM1",
		Baud: 9600,
	})
	if err != nil {
		t.Fatal(err)
	}

	connection, err := Connect(
		"serial",
		serial.WithPort(port),
	)
	if err != nil {
		t.Fatal(err)
	}
	defer connection.Close()

	for event := range <-connection.Events() {
		t.Log(event)
	}

}
