package service

import (
	"bytes"
	"testing"
)

func TestCommandStatusString(t *testing.T) {
	theory := func(m CommandStatus, expected string) {
		if m.String() != expected {
			t.Errorf("Expected %v recieved %v", m, expected)
		}
	}

	theory(CommandStatus(CommandFailure), "Failure")
	theory(CommandStatus(CommandSuccess), "Success")
}

func TestCommandStatusMarshalJSON(t *testing.T) {
	theory := func(m CommandStatus, expected string) {
		a, _ := m.MarshalJSON()
		e := []byte(expected)

		if bytes.Equal(a, e) {
			t.Errorf("Expected %v recieved %v", a, e)
		}
	}

	theory(CommandStatus(MuteOff), "Failure")
	theory(CommandStatus(MuteOn), "Success")
}

func TestCommandStatusUnMarshalJSON(t *testing.T) {
	theory := func(actual string, expected CommandStatus) {
		m := CommandStatus(0)

		a := []byte(actual)
		_ = m.UnmarshalJSON(a)

		if m != expected {
			t.Errorf("Expected %v recieved %v", expected, m)
		}
	}

	theory("\"Failure\"", CommandStatus(CommandFailure))
	theory("\"Success\"", CommandStatus(CommandSuccess))
}
