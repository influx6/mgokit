package common

import "sync"

// EventBroadcastSubscriber defines a interface that which is used to subscribe specifically for
// events  EventBroadcast type.
type EventBroadcastSubscriber interface {
	Receive(EventBroadcast)
}

//=========================================================================================================

// EventBroadcastHandler defines a structure type which implements the
// EventBroadcastSubscriber interface and the EventDistributor interface.
type EventBroadcastHandler struct {
	handle func(EventBroadcast)
}

// NewEventBroadcastHandler returns a new instance of a EventBroadcastHandler.
func NewEventBroadcastHandler(fn func(EventBroadcast)) *EventBroadcastHandler {
	return &EventBroadcastHandler{
		handle: fn,
	}
}

// Receive takes the giving value and execute it against the underline handler.
func (sn *EventBroadcastHandler) Receive(elem EventBroadcast) {
	sn.handle(elem)
}

// Handle takes the giving value and asserts the expected value to match the
// EventBroadcast type then passes it to the Receive method.
func (sn *EventBroadcastHandler) Handle(receive interface{}) {
	if elem, ok := receive.(EventBroadcast); ok {
		sn.Receive(elem)
	}
}

//=========================================================================================================

// EventBroadcastNotification defines a structure type which must be used to
// receive EventBroadcast type has a event.
type EventBroadcastNotification struct {
	sml        sync.Mutex
	subs       []EventBroadcastSubscriber
	validation func(EventBroadcast) bool
	register   map[EventBroadcastSubscriber]int
}

// NewEventBroadcastNotificationWith returns a new instance of EventBroadcastNotification.
func NewEventBroadcastNotificationWith(validation func(EventBroadcast) bool) *EventBroadcastNotification {
	var elem EventBroadcastNotification

	elem.validation = validation
	elem.register = make(map[EventBroadcastSubscriber]int, 0)

	return &elem
}

// NewEventBroadcastNotification returns a new instance of NewEventBroadcastNotification.
func NewEventBroadcastNotification() *EventBroadcastNotification {
	var elem EventBroadcastNotification
	elem.register = make(map[EventBroadcastSubscriber]int, 0)

	return &elem
}

// UnNotify removes the given subscriber from the notification's list if found from future events.
func (sn *EventBroadcastNotification) UnNotify(sub EventBroadcastSubscriber) {
	sn.do(func() {
		index, ok := sn.register[sub]
		if !ok {
			return
		}

		sn.subs = append(sn.subs[:index], sn.subs[index+1:]...)
	})
}

// Notify adds the given subscriber into the notification list and will await an update of
// a new event of the given EventBroadcast type.
func (sn *EventBroadcastNotification) Notify(sub EventBroadcastSubscriber) {
	sn.do(func() {
		sn.register[sub] = len(sn.subs)
		sn.subs = append(sn.subs, sub)
	})
}

// Handle takes the giving value and asserts the expected value to be of
// the type and pass on to it's underline subscribers else ignoring the event.
func (sn *EventBroadcastNotification) Handle(elem interface{}) {
	if elemEvent, ok := elem.(EventBroadcast); ok {
		if sn.validation != nil && sn.validation(elemEvent) {
			sn.do(func() {
				for _, sub := range sn.subs {
					sub.Receive(elemEvent)
				}
			})

			return
		}

		sn.do(func() {
			for _, sub := range sn.subs {
				sub.Receive(elemEvent)
			}
		})
	}
}

// do performs action with the mutex locked and unlocked appropriately, ensuring safe
// concurrent access.
func (sn *EventBroadcastNotification) do(fn func()) {
	if fn == nil {
		return
	}

	sn.sml.Lock()
	defer sn.sml.Unlock()

	fn()
}
