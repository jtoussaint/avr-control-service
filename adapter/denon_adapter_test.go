package adapter

import (
	"errors"
	"net"
	"testing"
	"time"
)

type connMock struct {
	expectedCommand    string
	expectErr          bool
	expectedMuteStatus MuteStatus
	muteStatus         MuteStatus
	muteStatusString   string
	t                  *testing.T
}

func (m connMock) Read(b []byte) (n int, err error) {
	copy(b[:], m.muteStatusString)

	return len(m.muteStatusString), nil
}

func (m connMock) Write(b []byte) (n int, err error) {
	s := string(b)

	if s != m.expectedCommand {
		m.t.Errorf("Expected %v recieved %v", m.expectedCommand, s)
	}

	return 0, nil
}

func (m connMock) Close() error {
	return nil
}

func (m connMock) LocalAddr() net.Addr {
	return nil
}

func (m connMock) RemoteAddr() net.Addr {
	return nil
}

func (m connMock) SetDeadline(t time.Time) error {
	return nil
}

func (m connMock) SetReadDeadline(t time.Time) error {
	return nil
}

func (m connMock) SetWriteDeadline(t time.Time) error {
	return nil
}

type readStringMock struct {
	actualString   string
	actualError    error
	expectedError  error
	expectedString string
}

func (m readStringMock) Read(b []byte) (n int, err error) {
	copy(b[:], m.actualString)

	return len(m.actualString), m.actualError
}

func TestReadMuteStatus(t *testing.T) {
	theory := func(m connMock) {
		a := DenonAdapter{AVRConnection: m}

		avr := AVR{}

		a.ReadMuteStatus(&avr)

		if avr.MuteStatus != m.expectedMuteStatus {
			m.t.Errorf("Expected %v recieved %v", m.expectedMuteStatus, avr.MuteStatus)
		}
	}

	theory(connMock{
		expectedCommand:    "MU?",
		expectedMuteStatus: MuteOff,
		muteStatusString:   "MUOFF",
		t:                  t,
	})

	theory(connMock{
		expectedCommand:    "MU?",
		expectedMuteStatus: MuteOn,
		muteStatusString:   "MUON",
		t:                  t,
	})

	theory(connMock{
		expectedCommand:    "MU?",
		expectedMuteStatus: MuteUnknown,
		muteStatusString:   "FOO",
		t:                  t,
	})
}

func TestSendMuteCommand(t *testing.T) {
	theory := func(m connMock) {
		a := DenonAdapter{AVRConnection: m}

		err := a.SendMuteCommand(m.muteStatus)

		if !m.expectErr && err != nil {
			m.t.Errorf("Recieved err when we did not expect one %v", err)
		}
	}

	theory(connMock{
		expectedCommand: "MUOFF",
		muteStatus:      MuteOff,
		t:               t,
	})

	theory(connMock{
		expectedCommand: "MUON",
		muteStatus:      MuteOn,
		t:               t,
	})

	theory(connMock{
		expectedCommand: "MUON",
		expectErr:       true,
		muteStatus:      MuteUnknown,
		t:               t,
	})
}

func TestReadString(t *testing.T) {
	theory := func(m readStringMock) {
		s, err := ReadString(m)
		if err != nil && m.expectedError == nil {
			t.Errorf("Did not expect an error, but recieved %v", err)
		}

		if err != nil && s != "" {
			t.Errorf("Expected an empty string when returning an error.")
		}

		if s != m.expectedString {
			t.Errorf("Expected %v, recieved %v", m.expectedString, s)
		}
	}

	theory(readStringMock{
		actualString:   "foo",
		expectedString: "foo",
	})

	theory(readStringMock{
		actualError:   errors.New("Test error"),
		expectedError: errors.New("Test error"),
	})
}
