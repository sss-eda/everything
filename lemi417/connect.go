package lemi417

// Connection TODO
type Connection interface {
	ReadConfig() error
	ReadTime() error
	Events() <-chan Event
	Close() error
}
