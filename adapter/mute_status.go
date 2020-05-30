package adapter

import (
	"bytes"
	"encoding/json"
)

//
// MuteStatus is a typed in to describe the state of the AVR mute
//
type MuteStatus int

const (
	//
	// MuteUnknown means we don't know the state of the AVR mute status
	//
	MuteUnknown = iota

	//
	// MuteOff means that AVR is not muted
	//
	MuteOff = iota

	//
	// MuteOn means that the AVR is not muted
	//
	MuteOn = iota
)

//
// Creates a string from a mute status
//
func (s MuteStatus) String() string {
	return toMuteStatusString[s]
}

//
// A map of mute statuses to strings
///
var toMuteStatusString = map[MuteStatus]string{
	MuteOff:     "Off",
	MuteOn:      "On",
	MuteUnknown: "Unknown",
}

//
// A map of strings to mute statuses
//
var toMuteStatusID = map[string]MuteStatus{
	"Off":     MuteOff,
	"On":      MuteOn,
	"Unknown": MuteUnknown,
}

//
// MarshalJSON marshals the enum as a quoted json string
//
func (s MuteStatus) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toMuteStatusString[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

//
// UnmarshalJSON unmashals a quoted json string to the enum value
//
func (s *MuteStatus) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value, 'MuteUnknown' in this case.
	*s = toMuteStatusID[j]
	return nil
}
