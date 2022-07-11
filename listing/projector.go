package listing

import "log"

type Projector struct {
	instruments InstrumentRepository
	commands    CommandRepository
	events      EventRepository
}

func (projector *Projector) onInstrumentRegistered(event instrumentation.InstrumentRegisteredEvent) {
	instrument, err := projector.instruments.GetInstrumentByID(event.InstrumentID)
	if err != nil {
		log.Println(err)
		return
	}

	instrument.commands = append(instrument.commands, event.Command)
}

func (projector *Projector) onCommandSentToInstrument(event instrumentation.CommandSentToInstrumentEvent) {
	instrument, err := projector.instruments.GetInstrumentByID(event.InstrumentID)
	if err != nil {
		log.Println(err)
		return
	}

	instrument.commands = append(instrument.commands, event.Command)
}
