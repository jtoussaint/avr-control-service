package adapter

import (
	"bytes"
	"testing"
)

func TestMuteStatusString(t *testing.T) {
	theory := func(m MuteStatus, expected string) {
		if m.String() != expected {
			t.Errorf("Expected %v recieved %v", m, expected)
		}
	}

	theory(MuteStatus(MuteOff), "Off")
	theory(MuteStatus(MuteOn), "On")
	theory(MuteStatus(MuteUnknown), "Unknown")
}

func TestMuteStatusMarshalJSON(t *testing.T) {
	theory := func(m MuteStatus, expected string) {
		a, _ := m.MarshalJSON()
		e := []byte(expected)

		if bytes.Equal(a, e) {
			t.Errorf("Expected %v recieved %v", a, e)
		}
	}

	theory(MuteStatus(MuteOff), "Off")
	theory(MuteStatus(MuteOn), "On")
	theory(MuteStatus(MuteUnknown), "Unknown")
}

func TestMuteStatusUnMarshalJSON(t *testing.T) {
	theory := func(actual string, expected MuteStatus) {
		m := MuteStatus(0)

		a := []byte(actual)
		_ = m.UnmarshalJSON(a)

		if m != expected {
			t.Errorf("Expected %v recieved %v", expected, m)
		}
	}

	theory("\"Off\"", MuteStatus(MuteOff))
	theory("\"On\"", MuteStatus(MuteOn))
	theory("\"Unknown\"", MuteStatus(MuteUnknown))
}
