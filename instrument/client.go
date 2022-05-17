package instrument

// Client - The interface an instrument expects to be available.
type Client interface {
	Register(ID) error
}
