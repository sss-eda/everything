package lemi025

type ErrUnknownConnectionKind struct{}

func (err ErrUnknownConnectionKind) Error() string {
	return "Unknown connection kind"
}
