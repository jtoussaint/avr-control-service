package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jtoussaint/avr-control/adapter"
)

type mockAdapter struct {
	commandResponse        interface{}
	expectedAVR            adapter.AVR
	expectedHTTPStatusCode int
	requestBody            string
	successfulDial         bool
}

func (m *mockAdapter) Close() {
}

func (m *mockAdapter) Dial() (ok bool, err error) {
	if !m.successfulDial {
		return false, errors.New("dail error")
	}

	return true, nil
}

func (m *mockAdapter) ReadAVR() (avr adapter.AVR) {
	return m.expectedAVR
}

func (m *mockAdapter) ReadMuteStatus(avr *adapter.AVR) (err error) {
	return nil
}

func (m *mockAdapter) SendMuteCommand(s adapter.MuteStatus) (err error) {
	return nil
}

func TestHealthzRequestHandler(t *testing.T) {
	theory := func(m *mockAdapter) {
		f := AdapterFactory{
			AVRAdapter: m,
		}

		req, err := http.NewRequest("GET", "/healthz", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(f.HealthzRequestHandler)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != m.expectedHTTPStatusCode {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, m.expectedHTTPStatusCode)
		}
	}

	theory(&mockAdapter{
		expectedHTTPStatusCode: http.StatusOK,
		successfulDial:         true,
	})

	theory(&mockAdapter{
		expectedHTTPStatusCode: http.StatusInternalServerError,
		successfulDial:         false,
	})
}

func TestMuteCommandRequestHandler(t *testing.T) {
	theory := func(m *mockAdapter, body string) {
		f := AdapterFactory{
			AVRAdapter: m,
		}

		req, err := http.NewRequest("PUT", "/mute", strings.NewReader(body))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(f.MuteCommandRequestHandler)

		handler.ServeHTTP(rr, req)
		status := rr.Code

		if status != m.expectedHTTPStatusCode {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, m.expectedHTTPStatusCode)
		}

		// quit early if we are in a
		if status != http.StatusOK {
			return
		}

		bodyBytes := bytes.Buffer{}
		json.NewEncoder(&bodyBytes).Encode(m.commandResponse)
		expectedBody := bodyBytes.String()
		if rr.Body.String() != expectedBody {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expectedBody)
		}
	}

	theory(&mockAdapter{
		expectedHTTPStatusCode: http.StatusBadRequest,
		successfulDial:         false,
	}, "foo")

	theory(&mockAdapter{
		expectedHTTPStatusCode: http.StatusInternalServerError,
		successfulDial:         false,
	}, "{}")

	theory(&mockAdapter{
		commandResponse:        MuteCommmandResponse{Status: CommandSuccess},
		expectedAVR:            adapter.AVR{MuteStatus: adapter.MuteOff},
		expectedHTTPStatusCode: http.StatusOK,
		successfulDial:         true,
	}, "{}")
}

func TestStatusRequestHandler(t *testing.T) {
	theory := func(m *mockAdapter) {
		f := AdapterFactory{
			AVRAdapter: m,
		}

		req, err := http.NewRequest("GET", "/", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(f.StatusRequestHandler)

		handler.ServeHTTP(rr, req)
		status := rr.Code

		if status != m.expectedHTTPStatusCode {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, m.expectedHTTPStatusCode)
		}

		// quit early if we are in a
		if status != http.StatusOK {
			return
		}

		bodyBytes := bytes.Buffer{}
		json.NewEncoder(&bodyBytes).Encode(m.expectedAVR)
		expectedBody := bodyBytes.String()
		if rr.Body.String() != expectedBody {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expectedBody)
		}
	}

	theory(&mockAdapter{
		expectedHTTPStatusCode: http.StatusInternalServerError,
		successfulDial:         false,
	})

	theory(&mockAdapter{
		expectedAVR:            adapter.AVR{MuteStatus: adapter.MuteOff},
		expectedHTTPStatusCode: http.StatusOK,
		successfulDial:         true,
	})
}
