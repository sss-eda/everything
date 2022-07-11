package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/sss-eda/everything/instrumentation"
)

func Instruments(
	registerInstrument instrumentation.RegisterInstrument,
	getAllInstruments instrumentation.GetAllInstruments,
) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getAllInstruments(w, r)
		case http.MethodPost:
			registerInstrument(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			log.Println("unsupported method")
		}
	}
}

func RegisterInstrument(
	registerInstrumentService instrumentation.RegisterInstrument,
) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		request := instrumentation.RegisterInstrumentRequest{}
		err := decoder.Decode(&request)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "unable to parse request: %v", err)
			return
		}
		registerInstrumentService(w, request)
	}
}
