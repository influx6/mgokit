package common

// EventBroadcast defines a struct which gets published for the events.
//@notification:event
type EventBroadcast struct {
	EventName string      `json:"event"`
	EventID   string      `json:"event_id"`
	Event     EventObject `json:"event_object"`
}

// Deliver will deliver the giving events into the appropriate pipeline for
// subscribers.
func (sn *EventBroadcastHandler) Deliver(name, id string, receive EventObject) {
	sn.Handle(EventBroadcast{EventName: name, EventID: id, Event: receive})
}

// EventObject defines a interface for the basic methods which events needs
// expose.
type EventObject interface {
	RemoveEvent()
	Underlying() interface{}
}
