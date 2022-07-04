package registration

type Service interface {
	RegisterAggregate(AggregateType string) error
}
