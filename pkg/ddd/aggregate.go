package ddd

type Aggregate interface {
	Root() Entity
	Version() uint64
}

type BaseAggregate struct {
	root Entity
}

func NewAggregate(root Entity) *BaseAggregate {
	return &BaseAggregate{
		root: root,
	}
}

func (aggregate *BaseAggregate) Root() Entity {
	return aggregate.root
}

func (aggregate *BaseAggregate) Version() uint64 {
	return 0
}
