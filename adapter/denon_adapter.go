package adapter

import (
	"fmt"
	"io"
	"net"
	"strings"
	"time"
)

//
// AVRAdapter represents a connection to a stereo / avr
//
type AVRAdapter interface {

	//
	// Close shuts down the underlying connection
	//
	Close()

	//
	// Dial creates a connection to the AVR and sets the read timeout to a default
	// value of 5 seconds
	//
	Dial() (ok bool, err error)

	//
	// ReadAVR reads the current state of the AVR
	//
	ReadAVR() (avr AVR)

	//
	// Determines the current mute status from the AVR
	//
	ReadMuteStatus(avr *AVR) (err error)

	//
	// Sends a mute command to the AVR
	//
	SendMuteCommand(m MuteStatus) (err error)
}

//
// DenonAdapter represents a connect to an AVR
//
type DenonAdapter struct {
	//
	// A connection to the avr
	//
	AVRConnection net.Conn

	//
	// The IP address of the host
	//
	Host string

	//
	// The port to connect to
	//
	Port int
}

//
// Close shuts down the underlying connection
//
func (d *DenonAdapter) Close() {
	if d.AVRConnection != nil {
		d.AVRConnection.Close()
	}
}

//
// Dial creates a connection to the AVR and sets the read timeout to a default
// value of 5 seconds
//
func (d *DenonAdapter) Dial() (ok bool, err error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%v:%v", d.Host, d.Port))
	if err != nil {
		return false, err
	}

	// set SetReadDeadline
	err = conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	if err != nil {
		// Unable to set a read deadline
		return false, fmt.Errorf("SetReadDeadline failed : %v", err)
	}

	d.AVRConnection = conn

	return ok, nil
}

//
// ReadString reads a string from the provided reader, trimming all whitespace
//
func ReadString(r io.Reader) (s string, err error) {
	recvBuf := make([]byte, 1024)

	n, err := r.Read(recvBuf[:]) // recv data
	if err != nil {
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			// time out
			return "", fmt.Errorf("read timeout : %v", err)
		}

		// some error else, do something else, for example create new conn
		return "", fmt.Errorf("read error : %v", err)
	}

	trimmed := make([]byte, n)
	copy(trimmed[:], recvBuf[0:n])

	resp := string(trimmed)
	resp = strings.TrimSpace(resp)

	return resp, nil
}

//
// ReadAVR reads the current state of the AVR
//
func (d *DenonAdapter) ReadAVR() (avr AVR) {
	avr.Name = "Denon"

	d.ReadMuteStatus(&avr)

	return avr
}

//
// ReadMuteStatus determines the current mute status from the AVR
//
func (d *DenonAdapter) ReadMuteStatus(avr *AVR) (err error) {
	fmt.Fprintf(d.AVRConnection, "MU?")

	status, err := ReadString(d.AVRConnection)
	if err != nil {
		return fmt.Errorf("error reading mute status : %v", err)
	}

	switch status {
	case "MUOFF":
		avr.MuteStatus = MuteOff
	case "MUON":
		avr.MuteStatus = MuteOn
	default:
		avr.MuteStatus = MuteUnknown
	}

	return nil
}

//
// SendMuteCommand sends a mute command to the AVR
//
func (d *DenonAdapter) SendMuteCommand(m MuteStatus) (err error) {

	switch m {
	case MuteOff:
		fmt.Fprintf(d.AVRConnection, "MUOFF")

		return nil
	case MuteOn:
		fmt.Fprintf(d.AVRConnection, "MUON")

		return nil
	default:
		return fmt.Errorf("Unkown Mute Status %v", m)
	}
}
