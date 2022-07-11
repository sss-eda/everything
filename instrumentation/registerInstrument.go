package instrumentation

type RegisterInstrument func(RegisterInstrumentResponseWriter, *RegisterInstrumentRequest)

type RegisterInstrumentResponse struct{}

type RegisterInstrumentResponseWriter interface {
	Write(RegisterInstrumentResponse) error
}

type RegisterInstrumentRequest struct {
	ID string
}
