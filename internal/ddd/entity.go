package ddd

type EntityID string

type Entity interface {
	ID() EntityID
}

type BaseEntity struct {
	id EntityID
}

func (entity BaseEntity) ID() EntityID {
	return entity.id
}
