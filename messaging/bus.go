package messaging

import "github.com/sss-eda/everything/internal/ddd"

type Bus interface {
	Subscribe(ddd.Domain, ddd.ID) (<-chan Message, <-chan error, error)
	Publish(ddd.Domain, ddd.ID, Message) error
}
