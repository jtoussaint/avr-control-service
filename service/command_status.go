package service

import (
	"bytes"
	"encoding/json"
)

//
// CommandStatus is a typed in to describe the state of the AVR command
//
type CommandStatus int

const (
	//
	// CommandFailure means the command failed to expecute
	//
	CommandFailure = iota

	//
	// CommandSuccess means that the command successeded
	//
	CommandSuccess = iota
)

//
// Creates a string from a command status
//
func (s CommandStatus) String() string {
	return toCommandStatusString[s]
}

//
// A map of command statuses to strings
///
var toCommandStatusString = map[CommandStatus]string{
	CommandFailure: "Failure",
	CommandSuccess: "Success",
}

//
// A map of strings to command statuses
//
var toCommandStatusID = map[string]CommandStatus{
	"Failure": CommandFailure,
	"Success": CommandSuccess,
}

//
// MarshalJSON marshals the enum as a quoted json string
//
func (s CommandStatus) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toCommandStatusString[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

//
// UnmarshalJSON unmashals a quoted json string to the enum value
//
func (s *CommandStatus) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value, 'commandUnknown' in this case.
	*s = toCommandStatusID[j]
	return nil
}
