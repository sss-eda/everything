package websocket

import (
	"encoding/json"

	"github.com/gorilla/websocket"
	"github.com/sss-eda/everything/pkg/es"
)

type EventSourcingClient struct {
	ws *websocket.Conn
}

func (client *EventSourcingClient) Publish(event es.Event) error {
	return client.ws.WriteJSON(event)
}

func (client *EventSourcingClient) Subscribe(channel chan<- es.Event) error {
	for {
		_, message, err := client.ws.ReadMessage()
		if err != nil {
			return err
		}

		event := es.Event{}
		err = json.Unmarshal(message, &event)
		if err != nil {
			return err
		}

		channel <- event
	}
}

func (client *EventSourcingClient) Close() error {
	return client.ws.Close()
}
