package adapter

//
// An AVR represents the current state of the stero
//
type AVR struct {
	Name       string     `json:"name"`
	Host       string     `json:"host"`
	MuteStatus MuteStatus `json:"muteStatus"`
}

//
// Mute will tell the AVR to mute it's self if there is a change
// in status
//
func (a *AVR) Mute(m MuteStatus) (ok bool) {
	if a.MuteStatus == m {
		return false
	}

	a.MuteStatus = m
	return true
}
