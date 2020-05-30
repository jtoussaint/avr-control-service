package service

import (
	"errors"
	"testing"
)

type mockAVRAdapter struct {
	calledSendMuteCommandSpy        bool
	currentMuteStatus               MuteStatus
	expectError                     bool
	expectedStatus                  CommandStatus
	newMuteStatus                   MuteStatus
	shouldHaveCalledSendMuteCommand bool
}

func (m *mockAVRAdapter) Close() {

}

func (m *mockAVRAdapter) Dial() (ok bool, err error) {
	if m.expectError {
		return false, errors.New("Error trying to dial")
	}

	return true, nil
}

func (m *mockAVRAdapter) ReadAVR() (avr AVR) {
	return AVR{
		MuteStatus: m.currentMuteStatus,
	}
}

func (m *mockAVRAdapter) SendMuteCommand(s MuteStatus) (err error) {
	m.calledSendMuteCommandSpy = true

	return nil
}

func (m *mockAVRAdapter) ReadMuteStatus(avr *AVR) (err error) {
	return nil
}

func TestHandle(t *testing.T) {
	theory := func(m mockAVRAdapter) {
		c := MuteCommand{AVRAdapter: &m, NewMuteStatus: m.newMuteStatus}

		r, err := c.Handle()

		if !m.expectError && err != nil {
			t.Errorf("Did not expect an error, but recieved %v", err)
			return
		}

		if m.shouldHaveCalledSendMuteCommand && !m.calledSendMuteCommandSpy {
			t.Errorf("Did not call SendMuteCommand")
		}

		if r.Status != m.expectedStatus {
			t.Errorf("Expected success, recieved %v", r.Status)
		}
	}

	theory(mockAVRAdapter{
		expectError:    true,
		expectedStatus: CommandFailure,
	})

	theory(mockAVRAdapter{
		currentMuteStatus:               MuteOff,
		expectError:                     false,
		expectedStatus:                  CommandSuccess,
		newMuteStatus:                   MuteOn,
		shouldHaveCalledSendMuteCommand: true,
	})

	theory(mockAVRAdapter{
		currentMuteStatus:               MuteOff,
		expectError:                     false,
		expectedStatus:                  CommandSuccess,
		newMuteStatus:                   MuteOff,
		shouldHaveCalledSendMuteCommand: false,
	})
}
