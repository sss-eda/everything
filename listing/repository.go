package listing

type InstrumentRepository interface {
	CreateInstrument(string) error
	GetInstrumentByID(string) (Instrument, error)
	GetAllInstruments() []Instrument
}

type CommandRepository interface {
	GetAllInstrumentCommands(string) []Command
}

type EventRepository interface {
	GetAllInstrumentEvents(string) []Event
}
