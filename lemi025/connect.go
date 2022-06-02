package lemi025

// Connection TODO
type Connection interface {
	ReadConfig() error
	ReadTime() error
	SetTime(Time) error
	SetCoefficients1(Coefficients1) error
	ReadCoefficients1() error
	SetCoefficients2(Coefficients2) error
	ReadCoefficients2() error
	Events() <-chan Event
	Close() error
}
