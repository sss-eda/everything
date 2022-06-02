package everything

// InstrumentationServerAPI TODO
type InstrumentationServerAPI interface {
	Events(InstrumentID) <-chan Event
	SendCommand(InstrumentID, Event) error
}

type InstrumentationClientAPI interface {
}
