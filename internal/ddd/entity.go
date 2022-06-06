package ddd

type EntityID string

type Identifiable interface {
	ID() EntityID
	IsEqualTo(Identifiable) bool
}

type Entity interface {
	Identifiable
}
