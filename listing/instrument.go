package listing

type Instrument struct {
	commands []Command
	events   []Event
}

func (instrument *Instrument) GetAllCommands() ([]Command, error) {
	return instrument.commands, nil
}

func (instrument *Instrument) GetAllEvents() ([]Event, error) {
	return instrument.events, nil
}
