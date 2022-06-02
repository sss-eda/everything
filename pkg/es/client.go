package es

type Client interface {
	Publish(Event) error
	Subscribe(chan<- Event) error
	Close() error
}
