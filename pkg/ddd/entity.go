package ddd

type EntityID string

// Entity TODO
type Entity interface {
	ID() EntityID
}
