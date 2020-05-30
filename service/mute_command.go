package service

import (
	"fmt"

	"github.com/jtoussaint/avr-control/adapter"
)

//
// MuteCommand tells the AVR to mute
//
type MuteCommand struct {
	AVRAdapter    adapter.AVRAdapter
	NewMuteStatus adapter.MuteStatus `json:"newMuteStatus"`
}

//
// MuteCommmandResponse details the response of the mute command
//
type MuteCommmandResponse struct {
	Status CommandStatus `json:"status"`
}

//
// Handle performs the mute operation
//
func (c *MuteCommand) Handle() (r MuteCommmandResponse, err error) {
	if c.AVRAdapter == nil {
		return MuteCommmandResponse{Status: CommandFailure}, fmt.Errorf("Expected an AVR Adapter")
	}

	_, err = c.AVRAdapter.Dial()
	if err != nil {
		return MuteCommmandResponse{Status: CommandFailure}, fmt.Errorf("Unable to connect to avr %v", err)
	}

	defer c.AVRAdapter.Close()

	avr := c.AVRAdapter.ReadAVR()
	ok := avr.Mute(c.NewMuteStatus)

	if ok {
		c.AVRAdapter.SendMuteCommand(avr.MuteStatus)
	}

	return MuteCommmandResponse{Status: CommandSuccess}, nil
}
