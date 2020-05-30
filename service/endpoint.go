package service

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jtoussaint/avr-control/adapter"
)

//
// AdapterFactory provides access to backend systems
//
type AdapterFactory struct {
	AVRAdapter adapter.AVRAdapter
}

//
// HealthzRequestHandler checks for service health
//
func (f AdapterFactory) HealthzRequestHandler(w http.ResponseWriter, r *http.Request) {
	_, err := f.AVRAdapter.Dial()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer f.AVRAdapter.Close()
}

//
// MuteCommandRequestHandler to mute or unmute the AVR
//
func (f AdapterFactory) MuteCommandRequestHandler(w http.ResponseWriter, r *http.Request) {
	c := MuteCommand{AVRAdapter: f.AVRAdapter}

	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response, err := c.Handle()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(response)
}

//
// StatusRequestHandler returns the current status of the AVR
//
func (f AdapterFactory) StatusRequestHandler(w http.ResponseWriter, r *http.Request) {
	_, err := f.AVRAdapter.Dial()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer f.AVRAdapter.Close()

	avr := f.AVRAdapter.ReadAVR()
	json.NewEncoder(w).Encode(avr)
}
