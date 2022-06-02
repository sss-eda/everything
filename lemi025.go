package everything

// Lemi025 TODO
type Lemi025 interface {
	ReadConfig() error
	ReadTime() error
	SetTime(Time) error
	Events() <-chan Event
	Close() error
}
