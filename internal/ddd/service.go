package ddd

type Service struct {
	AggregateTypes []AggregateType
}

func NewService(aggregateTypes ...AggregateType) (*Service, error) {
	return &Service{aggregateTypes...}
}
