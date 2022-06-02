package lemi025

import "github.com/sss-eda/everything/pkg/ddd"

type Instrument struct {
	id string
}

func NewInstrument(id string) *Instrument {
	return &Instrument{
		id: id,
	}
}

func (instrument *Instrument) ID() ddd.EntityID {
	return ddd.EntityID(instrument.id)
}
