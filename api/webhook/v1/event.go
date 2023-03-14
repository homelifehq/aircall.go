package v1

type Event string

const ErrInvalidEvent = "invalid event"

const (
	UserCreated              Event = "user.created"
	UserDeleted              Event = "user.deleted"
	UserConnected            Event = "user.connected"
	UserDisconnected         Event = "user.disconnected"
	UserOpened               Event = "user.opened"
	UserClosed               Event = "user.closed"
	UserWutStart             Event = "user.wut_start"
	UserWutEnd               Event = "user.wut_end"
	NumberCreated            Event = "number.created"
	NumberDeleted            Event = "number.deleted"
	NumberOpened             Event = "number.opened"
	NumberClosed             Event = "number.closed"
	CallCreated              Event = "call.created"
	CallRingingOnAgent       Event = "call.ringing_on_agent"
	CallAgentDeclined        Event = "call.agent_declined"
	CallAnswered             Event = "call.answered"
	CallTransferred          Event = "call.transferred"
	CallUnsuccessfulTransfer Event = "call.unsuccessful_transfer"
	CallHungup               Event = "call.hungup"
	CallEnded                Event = "call.ended"
	CallVoicemailLeft        Event = "call.voicemail_left"
	CallAssigned             Event = "call.assigned"
	CallArchived             Event = "call.archived"
	CallTagged               Event = "call.tagged"
	CallUntagged             Event = "call.untagged"
	CallCommented            Event = "call.commented"
	CallHold                 Event = "call.hold"
	CallUnhold               Event = "call.unhold"
	ContactCreated           Event = "contact.created"
	ContactUpdated           Event = "contact.updated"
	ContactDeleted           Event = "contact.deleted"
)

// String returns the string representation of the event
func (e Event) String() string {
	return string(e)
}

// IsValid returns true if the event is valid
func (e Event) IsValid() bool {
	switch e {
	case UserCreated,
		UserDeleted,
		UserConnected,
		UserDisconnected,
		UserOpened,
		UserClosed,
		UserWutStart,
		UserWutEnd,
		NumberCreated,
		NumberDeleted,
		NumberOpened,
		NumberClosed,
		CallCreated,
		CallRingingOnAgent,
		CallAgentDeclined,
		CallAnswered,
		CallTransferred,
		CallUnsuccessfulTransfer,
		CallHungup,
		CallEnded,
		CallVoicemailLeft,
		CallAssigned,
		CallArchived,
		CallTagged,
		CallUntagged,
		CallCommented,
		CallHold,
		CallUnhold,
		ContactCreated,
		ContactUpdated,
		ContactDeleted:
		return true
	}

	return false
}
