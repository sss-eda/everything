package lemi025

import (
	"encoding/json"
	"log"

	"github.com/sss-eda/everything"
)

func Run(
	done <-chan any,
	inputStream <-chan []byte,
) <-chan []byte {
	outputStream := make(chan []byte)

	go func() {
		defer close(outputStream)

		var err error

		for {
			select {
			case <-done:
				return
			case input := <-inputStream:
				message := everything.Message{}
				err = json.Unmarshal(input, &message)
				if err != nil {
					log.Println(err)
					return
				}
				switch message.Kind {
				case everything.CommandMessageKind:
					command := Command{}
					err = json.Unmarshal(message.Payload, &command)
					if err != nil {
						log.Println(err)
						return
					}
					switch command.Kind {
					case ReadConfigCommandKind:
						readConfigCommand = ReadConfigCommand{}
						err = json.Unmarshal(command.Payload, &readConfigCommand)
						if err != nil {
							log.Println(err)
							return
						}
						readConfigCommand.Execute()
					case ReadTimeCommandKind:
					case SetTimeCommandKind:
					}
				}
			}
		}
	}()

	return outputStream
}
