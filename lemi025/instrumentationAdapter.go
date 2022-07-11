package lemi025

type Command struct {
	Kind    string
	Payload []byte
}

type Instrument struct {
	// station  *Station
	commands map[uint64]Command
}

// func (instrument *Instrument) SendCommand(
// 	ctx context.Context,
// 	command Command,
// ) error {
// 	switch command.Kind {
// 	case "read-config":
// 		return instrument.station.ReadConfig(command.Payload)
// 	case "read-time":
// 		return instrument.station.ReadTime(command.Payload)
// 	default:
// 		return fmt.Errorf("command kind not supported: %v", command.Kind)
// 	}
// }

// func (instrument *Instrument) OnCommandSent(
// 	ctx context.Context,
// 	event CommandSentEvent,
// ) {
// 	instrument.
// }

// type CommandSentEvent (BadExpr)
