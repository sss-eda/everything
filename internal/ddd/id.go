package ddd

type ID interface {
	Equal(ID) bool
	Less(ID) bool
}
