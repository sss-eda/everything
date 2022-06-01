package lemi025

import "io"

type ReadConfigCommand func(ReadConfigCommandPayload) error
type ReadConfigCommandPayload struct{}

func ReadConfig(writer io.Writer) ReadConfigCommand {
	return func(payload ReadConfigCommandPayload) error {
		_, err := writer.Write([]byte{0x3D, 0x30})
		return err
	}
}
