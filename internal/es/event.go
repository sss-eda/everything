package es

type Event struct {
	Version uint64
	Kind    string
	Payload any
}