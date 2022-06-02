package lemi025

import "github.com/sss-eda/everything/pkg/cqrs"

// Message TODO
type Message struct {
	Kind string `json:"kind"`
	Payload []byte `json:"payload"`
}

func Run(
	done <-chan any,
	input <-chan []byte,
) <-chan []byte {
	var err error
	output := make(chan []byte)

	go func() {
		defer close(output)
		
		for in := range input {
			select {
				case <-done:
					return
				default:
					message := Message{}
					err = json.Unmarshal(in, &message)
					if err != nil {
						log.Println(err)
						continue
					}
					switch message.Kind {
						case "event":
							event := es.Event{}
							err := json.Unmarshal(message.Payload, &event)
							if err != nil {
								log.Println(err)
								continue
							}

						case "command":
							command := cqrs.Command{}
							err := json.Unmarshal(message.Payload, &command)
							if err != nil {
								log.Println(err)
								continue
							}
					}

			}
		}
	}()

	go func() {
		defer close(newCommands)
		
		for command := range commands {

		}
	}()

	return newEvents, newCommands
}

multiply := func(
  done <-chan interface{},
  intStream <-chan int,
  multiplier int,
) <-chan int {
    multipliedStream := make(chan int)
    go func() {
        defer close(multipliedStream)
        for i := range intStream {
            select {
            case <-done:
                return
            case multipliedStream <- i*multiplier:
            }
        }
    }()
    return multipliedStream
}

add := func(
  done <-chan interface{},
  intStream <-chan int,
  additive int,
) <-chan int {
    addedStream := make(chan int)
    go func() {
        defer close(addedStream)
        for i := range intStream {
            select {
            case <-done:
                return
            case addedStream <- i+additive:
            }
        }
    }()
    return addedStream
}

done := make(chan interface{})
defer close(done)

intStream := generator(done, 1, 2, 3, 4)
pipeline := multiply(done, add(done, multiply(done, intStream, 2), 1), 2)

for v := range pipeline {
    fmt.Println(v)
}
