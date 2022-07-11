package instrumentation

type Controller interface {
	SendCommandToInstrument(Command) error
}
