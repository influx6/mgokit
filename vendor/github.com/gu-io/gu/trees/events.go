package trees

import (
	"fmt"
	"strings"

	"github.com/gu-io/gu/common"
)

// EventOptions defines a function type used to apply specific operations to a
// Event object.
type EventOptions func(*Event)

// EventType sets the type of event.
func EventType(eventType string) EventOptions {
	return func(ev *Event) {
		ev.Type = eventType
	}
}

// EventTarget sets the type of event.
func EventTarget(target string) EventOptions {
	return func(ev *Event) {
		ev.secTarget = target
	}
}

// StopImmediatePropagation sets the type of event.
func StopImmediatePropagation(state bool) EventOptions {
	return func(ev *Event) {
		ev.StopImmediatePropagation = state
	}
}

// StopPropagation sets the type of event.
func StopPropagation(state bool) EventOptions {
	return func(ev *Event) {
		ev.StopPropagation = state
	}
}

// UseCapture sets the type of event.
func UseCapture(state bool) EventOptions {
	return func(ev *Event) {
		ev.UseCapture = state
	}
}

// PreventDefault sets the type of event.
func PreventDefault(state bool) EventOptions {
	return func(ev *Event) {
		ev.PreventDefault = state
	}
}

// Event provide a meta registry for helps in registering events for dom markups
// which is translated to the nodes themselves
type Event struct {
	Type                     string
	PreventDefault           bool
	StopPropagation          bool
	UseCapture               bool
	StopImmediatePropagation bool
	Tree                     *Markup
	Remove                   common.Remover
	secTarget                string
}

// NewEvent returns a event object that allows registering events to eventlisteners.
func NewEvent(options ...EventOptions) *Event {
	evm := &Event{}

	for _, option := range options {
		if option == nil {
			continue
		}
		option(evm)
	}

	return evm
}

// Target returns the target of the giving event.
func (e *Event) Target() string {
	if e.Tree != nil {
		return e.Tree.EventID()
	}

	return e.secTarget
}

// EventJSON defines a struct which contains the giving events and
// and tree of the giving tree.
type EventJSON struct {
	ParentSelector           string `json:"ParentSelector"`
	EventSelector            string `json:"EventSelector"`
	EventName                string `json:"EventName"`
	Event                    string `json:"Event"`
	PreventDefault           bool   `json:"PreventDefault"`
	StopPropagation          bool   `json:"StopPropagation"`
	UseCapture               bool   `json:"UseCapture"`
	StopImmediatePropagation bool   `json:"StopImmediatePropagation"`
}

// EventJSON returns the event json structure which represent the giving event.
func (e *Event) EventJSON() EventJSON {
	return EventJSON{
		Event:                    e.Type,
		UseCapture:               e.UseCapture,
		EventName:                e.EventName(),
		EventSelector:            e.EventSelector(),
		ParentSelector:           e.ParentEventSelector(),
		PreventDefault:           e.PreventDefault,
		StopPropagation:          e.StopPropagation,
		StopImmediatePropagation: e.StopImmediatePropagation,
	}
}

// ParentEventSelector returns the parent selector for this events markup tree.
func (e *Event) ParentEventSelector() string {
	if e.Tree != nil {
		return e.Tree.IDSelector(true)
	}

	return ""
}

// EventSelector returns the selector for this events tree.
func (e *Event) EventSelector() string {
	if e.Tree != nil && e.secTarget == "" {
		return e.Tree.IDSelector(false)
	}

	return e.secTarget
}

// EventName returns the giving name of the event.
func (e *Event) EventName() string {
	eventName := strings.ToUpper(e.Type[:1]) + e.Type[1:]
	if strings.HasSuffix(eventName, "Event") {
		return eventName
	}

	return eventName + "Event"
}

// ID returns the uique event id string for the event.
func (e *Event) ID() string {
	return fmt.Sprintf("%s#%s", e.Target(), e.Type)
}

// Clone  returns a new Event object from this.
func (e *Event) Clone() *Event {
	return &Event{
		Type:                     e.Type,
		secTarget:                e.secTarget,
		PreventDefault:           e.PreventDefault,
		UseCapture:               e.UseCapture,
		StopPropagation:          e.StopPropagation,
		StopImmediatePropagation: e.StopImmediatePropagation,
	}
}

// Apply adds the event into the elements events lists
func (e *Event) Apply(ex *Markup) {
	if !ex.allowEvents {
		return
	}

	e.Tree = ex

	ex.AddEvent(*e)
}

// String returns the string representation of the giving event.
func (e *Event) String() string {
	return fmt.Sprintf("%#v", e.EventJSON())
}

//==============================================================================
