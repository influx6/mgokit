package trees

import "sync"

// Morpher defines an interface which morphs the giving markup based on
// its current internal state based on some internal condition.
type Morpher interface {
	Morph(*Markup) *Markup
}

// SwitchMorpher defines an interface which allows switching the behaviour of
// which determines what the markup provided gets morphed into. This allows a nice
// simple binary morph of input based on certain conditions.
type SwitchMorpher interface {
	Morpher
	On(interface{})
	Off(interface{})
}

// RemoveMorpher defines a morpher which sets the giving markup as removed.
type RemoveMorpher struct {
	wl     sync.RWMutex
	remove bool
}

// On switches the state of the morpher out of a remove to be true.
func (r *RemoveMorpher) On(m interface{}) {
	r.wl.Lock()
	r.remove = true
	r.wl.Unlock()
}

// Off switches the state of the morpher to set remove to be true.
func (r *RemoveMorpher) Off(m interface{}) {
	r.wl.Lock()
	r.remove = false
	r.wl.Unlock()
}

// Morph sets the markup received as removed.
func (r *RemoveMorpher) Morph(m *Markup) *Markup {
	r.wl.RLock()
	if r.remove {
		m.Remove()
		r.wl.RUnlock()

		return m
	}

	m.UnRemove()
	r.wl.RUnlock()
	return m
}

// HideMorpher defines a morpher which sets the giving markup as removed.
type HideMorpher struct {
	wl     sync.RWMutex
	hidden bool
}

// On switches the state of the morpher out of a hidden state to be true.
func (r *HideMorpher) On(m interface{}) {
	r.wl.Lock()
	r.hidden = true
	r.wl.Unlock()
}

// Off switches the state of the morpher to set hidden state to be true.
func (r *HideMorpher) Off(m interface{}) {
	r.wl.Lock()
	r.hidden = false
	r.wl.Unlock()
}

// Morph sets the markup received as removed.
func (r *HideMorpher) Morph(m *Markup) *Markup {
	r.wl.RLock()
	{
		if r.hidden {
			Hide.Mode(m)
		} else {
			Show.Mode(m)
		}
	}
	r.wl.RUnlock()

	return m
}

//==============================================================================
