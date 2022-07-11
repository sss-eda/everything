package instrumentation

type GetAllInstruments func(GetAllInstrumentsResponseWriter, *GetAllInstrumentsRequest)

type GetAllInstrumentsResponse struct{}

type GetAllInstrumentsResponseWriter interface {
	Write(RegisterInstrumentResponse) error
}

type GetAllInstrumentsRequest struct {
	ID string
}
