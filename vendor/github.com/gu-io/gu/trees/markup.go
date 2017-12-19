package trees

import (
	"fmt"
	"html/template"
	"strings"

	"github.com/gu-io/gu/trees/css"
	"github.com/russross/blackfriday"
)

// Markup represent a concrete implementation of a element node.
type Markup struct {
	ID              string
	removed         bool
	autoclose       bool
	allowEvents     bool
	allowChildren   bool
	allowStyles     bool
	allowAttributes bool

	uid           string
	hash          string
	tagname       string
	textContent   string
	idSelector    string
	textContentFn func(*Markup) string

	events   []Event
	children []*Markup
	styles   []Property
	attrs    []Property
	morphers []Morpher
	parent   *Markup
}

// NewText returns a new Text instance element
func NewText(txt string, dl ...interface{}) *Markup {
	em := NewMarkup("text", false)
	em.allowChildren = false
	em.allowAttributes = false
	em.allowStyles = false
	em.allowEvents = false

	if dl != nil && len(dl) != 0 {
		em.textContent = fmt.Sprintf(txt, dl...)
	} else {
		em.textContent = txt
	}

	return em
}

// MarkdownTemplate returns a markup generated from a markup down string
// which is built into a markup. If an error occured, it will be turned into
// an error tag with the contents of the error.
func MarkdownTemplate(tml string, bind interface{}) *Markup {
	processed, err := Templated(tml, bind, func(in string) string {
		return string(blackfriday.MarkdownCommon([]byte(in)))
	})

	// if error occured, return a <error> tag with error details.
	if err != nil {
		return ParseFirstOrMakeRoot("<error>" + err.Error() + "</error>")
	}

	return ParseFirstOrMakeRoot(processed)
}

// CSSStylesheet provides a function that takes style rules which returns a stylesheet embeded into
// the provided element parent and is built on the gu/css package which collects
// necessary details from its parent to only target where it gets mounted.
func CSSStylesheet(styles interface{}, bind interface{}, ext *css.Rule, plain bool) *Markup {
	var rs *css.Rule

	switch so := styles.(type) {
	case string:
		switch plain {
		case true:
			rs = css.Plain(so, ext)
		case false:
			rs = css.New(so, ext)
		}
		break
	case *css.Rule:
		rs = so
		rs.UseExtension(ext)
		break
	default:
		panic("Invalid Acceptable type: Only string or *css.Rule")
	}

	content := NewMarkup("style", false)
	content.allowChildren = false
	content.allowAttributes = false
	content.allowStyles = false
	content.allowEvents = false
	content.textContentFn = func(owner *Markup) string {
		sheet, err := rs.Stylesheet(bind, owner.IDSelector(true))
		if err != nil {
			return err.Error()
		}

		return sheet.String()
	}

	return content
}

//==============================================================================

// NewMarkup returns a new element instance giving the specified name which is
// used as a tag name.
func NewMarkup(tag string, autoClose bool) *Markup {
	return &Markup{
		allowChildren:   true,
		allowStyles:     true,
		allowAttributes: true,
		allowEvents:     true,
		uid:             RandString(8),
		hash:            RandString(10),
		autoclose:       autoClose,
		tagname:         strings.ToLower(strings.TrimSpace(tag)),
		attrs:           []Property{NewAttr("data-gen", "gu")},
	}
}

// Empty resets the elements children list as 0 length
func (e *Markup) Empty() {
	e.children = nil
	e.events = nil
	e.styles = nil
	e.morphers = nil
}

// MarkupJSON defines a struct which contains the giving events and
// and tree of the giving tree.
type MarkupJSON struct {
	TreeID string      `json:"TreeID"`
	Events []EventJSON `json:"Events"`
	Markup string      `json:"Markup"`
}

// TreeJSON returns the giving MarkupJSON for this giving markup and
// the events related to this markup.
func (e *Markup) TreeJSON() MarkupJSON {
	var mjson MarkupJSON
	mjson.TreeID = e.uid

	// SetMode sets the mode into which we wish to print, we want the
	// details of removed.
	data, _ := e.MarshalJSON()
	mjson.Markup = string(data)

	e.EachEvent(func(event *Event, _ *Markup) {
		mjson.Events = append(mjson.Events, event.EventJSON())
	})

	return mjson
}

// IDSelector returns the unique selector for the giving markup.
func (e *Markup) IDSelector(useParent bool) string {
	var parentName string

	if e.parent != nil && useParent {
		if e.parent.ID == "" {
			parentName = e.parent.tagname + "[uid='" + e.parent.uid + "']"
		} else {
			parentName = "#" + e.parent.ID
		}

		return parentName
	}

	if e.ID == "" {
		parentName = e.tagname + "[uid='" + e.uid + "']"
	} else {
		parentName = "#" + e.ID
	}

	return parentName
}

// MarshalJSON returns the html representation of the giving markup.
func (e *Markup) MarshalJSON() ([]byte, error) {
	return []byte(e.HTML()), nil
}

// UnmarshalJSON the provided data and adds the giving children into the giving root.
func (e *Markup) UnmarshalJSON(data []byte) error {
	parsed := ParseTree(string(data))

	if len(parsed) == 1 {
		item := parsed[0]
		e.attrs = item.attrs
		e.textContent = item.textContent
		e.textContentFn = item.textContentFn
		e.tagname = item.tagname
		e.styles = item.styles
		e.events = item.events
		e.allowStyles = item.allowStyles
		e.allowAttributes = item.allowAttributes
		e.allowEvents = item.allowEvents
		e.allowChildren = item.allowChildren
		e.ID = item.ID

		item = nil
		parsed = nil
		return nil
	}

	if e.tagname == "" {
		e.tagname = "div"
	}

	e.AddChild(parsed...)
	return nil
}

// EHTML returns the html string wrapped by a template.HTML type to avoid getting
// escaped by go templates. The returned html is rendered using the default
// SimpleElementWriter and represents the DOM of the giving element.
func (e *Markup) EHTML() template.HTML {
	return template.HTML(SimpleElementWriter.Print(e))
}

// HTML returns the html string representing the DOM of the giving element.
//The returned html is rendered using the default SimpleElementWriter.
func (e *Markup) HTML() string {
	return SimpleElementWriter.Print(e)
}

// AutoClosed returns true/false if this element uses a </> or a <></> tag convention
func (e *Markup) AutoClosed() bool {
	return e.autoclose
}

//==============================================================================

// Properties interface defines a type that has Attributes
type Properties interface {
	AddAttribute(Property)
	Attributes() []Property

	AddStyle(Property)
	Styles() []Property

	AddEvent(Event)
	Events() []Event
}

// AddEvent adds an event into the event list for this element.
func (e *Markup) AddEvent(ev Event) {
	e.events = append(e.events, ev)
}

// EachEvent iterates all events from this giving root down with all childrens
// allowing the callback to process the events has needed.
func (e *Markup) EachEvent(fn func(*Event, *Markup)) {
	for _, ev := range e.events {
		fn(&ev, e)
	}

	for _, ch := range e.children {
		ch.EachEvent(fn)
	}
}

// Events return the elements events
func (e *Markup) Events() []Event {
	return e.events
}

// Styles return the internal style list of the element
func (e *Markup) Styles() []Property {
	return e.styles
}

// AddStyle adds a property to the style property list.
func (e *Markup) AddStyle(p Property) {
	e.styles = append(e.styles, p)
}

// Attributes return the internal attribute list of the element
func (e *Markup) Attributes() []Property {
	return e.attrs
}

// AddAttribute adds a property to the attribute property list.
func (e *Markup) AddAttribute(p Property) {
	e.attrs = append(e.attrs, p)
}

//==============================================================================

// EventID returns the selector used for tagging events for a markup.
func (e *Markup) EventID() string {
	return fmt.Sprintf("%s[uid='%s']", strings.ToLower(e.Name()), e.UID())
}

// Name returns the tag name of the element
func (e *Markup) Name() string {
	return e.tagname
}

// UID returns the current uid of the Element
func (e *Markup) UID() string {
	return e.uid
}

// Hash returns the current hash of the Element
func (e *Markup) Hash() string {
	return e.hash
}

//==============================================================================

// Morphers exposes a method to allow adding morphers.
type Morphers interface {
	Morphers() []Morphers
	AddMorpher(...Morpher)
	ApplyMorphers() *Markup
}

// AddMorpher adds the provided morphers into the elements lists.
func (e *Markup) AddMorpher(m ...Morpher) {
	e.morphers = append(e.morphers, m...)
}

// Morphers returns the slice of Morphers used by this markup.
func (e *Markup) Morphers() []Morpher {
	return e.morphers
}

// ApplyMorphers calls all elemental morphers sequentially applying them to the
// element and passes the result as the input of the next morpher unless. If
// any morpher returns nil, then the element is reused again until all morphers
// are called.
func (e *Markup) ApplyMorphers() *Markup {
	for _, child := range e.children {
		child.ApplyMorphers()
	}

	for _, morpher := range e.morphers {
		morpher.Morph(e)
	}

	return e
}

//==============================================================================

// TextMarkup defines a interface for text based markup.
type TextMarkup interface {
	TextContent() string
}

// TextContent returns the elements text value by either running a text content
// function if provided else defaulting to the textContent field.
func (e *Markup) TextContent() string {
	if e.textContentFn != nil {
		return e.textContentFn(e)
	}

	return e.textContent
}

// Clean cleans out all internal markup marked as removable.
func (e *Markup) Clean() {
	for n, elm := range e.children {
		if elm.Removed() {
			copy(e.children[n:], e.children[n+1:])
			e.children = e.children[:len(e.children)-1]
		} else {
			elm.Clean()
		}
	}
}

// Remove sets the markup as removable and adds a 'NodeRemoved' attribute to it.
func (e *Markup) Remove() {
	if !e.Removed() {
		e.attrs = append(e.attrs, &Attribute{Name: "NodeRemoved", Value: ""})
		e.removed = true
	}
}

// UnRemove sets the markup as not to be removable.
func (e *Markup) UnRemove() {
	if !e.Removed() {
		return
	}

	e.removed = false

	for index, attr := range e.attrs {
		if name, _ := attr.Render(); name != "NodeRemoved" {
			continue
		}

		e.attrs = append(e.attrs[:index], e.attrs[1+index:]...)
		return
	}
}

// Removed returns true/false if the Element is marked removed
func (e *Markup) Removed() bool {
	return !!e.removed
}

// SwapUID swaps the uid of the internal Element.
func (e *Markup) SwapUID(uid string) {
	e.uid = uid
}

// SwapHash swaps the hash of the internal Element.
func (e *Markup) SwapHash(hash string) {
	e.hash = hash
}

// UpdateHash updates the Element hash value
func (e *Markup) UpdateHash() {
	e.hash = RandString(10)
}

// Reconcile takes a old markup and reconciles its uid and its children with
// these information,it returns a true/false telling the parent if the children
// swapped hashes.
// The reconcilation uses the order in which elements are added, if the order
// and element types are same then the uid are swapped, else it firsts checks the
// element type, but if not the same adds the old one into the new list as removed
// then continues the check. The system takes position of elements in the old and
// new as very important and I cant stress this enough, "Element Positioning" in
// the markup are very important, If a Anchor was the first element in the old
// render and the next pass returns a Div in the position for that Anchor in the
// new render, the old Anchor will be marked as removed and will be removed from
// the dom and ignored by the writers.
// When two elements position are same and their types are the same then a checkup
// process is done using the elements attributes, this is done to determine if the
// hash value of the new should be swapped with the old. We cant use style properties
// here because they are the most volatile of the set and will periodically be
// either changed and returned to normal values eg display: none to display: block
// and vise-versa, so only attributes are used in the check process.
func (e *Markup) Reconcile(em *Markup) bool {
	if e == em {
		return false
	}

	// are we reconciling the proper elements type ? if not skip (i.e different types cant reconcile eachother)]
	// TODO: decide if we should mark the markup as removed in this case as a catchall system
	if e.Name() != em.Name() {
		return false
	}

	em.Clean()

	//since the tagname are the same, swap uids
	// olduid := em.UID()
	e.SwapUID(em.UID())

	//since the tagname are the same and we have swapped uid, to determine who gets or keeps
	// its hash we will check the attributes against each other, but also the hash is dependent on the
	// children also, if the children observered there was a change
	oldHash := em.Hash()

	// if we have a special case for text element then we do things differently
	if e.Name() == "text" {
		if e.TextContent() == em.TextContent() {
			e.SwapHash(oldHash)
			return false
		}

		return true
	}

	newChildren := e.Children()
	oldChildren := em.Children()

	maxSize := len(newChildren)
	oldMaxSize := len(oldChildren)

	equalAttr := EqualAttributes(e, em)
	equalStyle := EqualStyles(e, em)

	// if the element had no children too, swap hash.
	if maxSize == 0 {
		if oldMaxSize > 0 {
			return true
		}

		if !equalAttr && !equalStyle {
			e.SwapHash(oldHash)
			return false
		}

		return true
	}

	var childChanged bool

	for n, och := range oldChildren {
		if maxSize > n {

			nch := newChildren[n]
			if nch.Name() != och.Name() {

				och.Remove()
				e.AddChild(och)
				childChanged = true
				continue
			}

			if nch.Reconcile(och) {
				childChanged = true
			}

			continue
		}

		och.Remove()
		e.AddChild(och)
		childChanged = true
	}

	if !childChanged && equalAttr && equalStyle {
		e.SwapHash(oldHash)
		return false
	}

	return true
}

// FirstChild returns the first child in the markup children list.
func (e *Markup) FirstChild() *Markup {
	return e.NthChild(0)
}

// LastChild returns the last child in the markup children list.
func (e *Markup) LastChild() *Markup {
	return e.NthChild(len(e.children) - 1)
}

// NthChild returns the giving child at the index position.
func (e *Markup) NthChild(index int) *Markup {
	childrenLen := len(e.children)

	if index >= childrenLen && index <= -1 {
		return nil
	}

	return e.children[index]
}

// AddChild adds a new markup as the children of this element
func (e *Markup) AddChild(child ...*Markup) {
	if !e.allowChildren {
		return
	}

	for _, ch := range child {
		if ch == e || ch == nil {
			continue
		}

		ch.parent = e
		e.children = append(e.children, ch)
	}
}

// EachChild iterates all children from this giving root down with all childrens
// allowing the callback to process the child has needed.
func (e *Markup) EachChild(fn func(*Markup)) {
	for _, ch := range e.children {
		fn(ch)

		ch.EachChild(fn)
	}
}

// Children returns the children list for the element
func (e *Markup) Children() []*Markup {
	return e.children
}

//==============================================================================

// Appliable define the interface specification for applying changes to elements elements in tree
type Appliable interface {
	Apply(*Markup)
}

// ApplyChildren applies the internal children of the markup to the provided parent.
func (e *Markup) ApplyChildren(em *Markup) {
	if em == nil {
		return
	}

	for _, child := range e.children {
		em.AddChild(child)
	}
}

//Apply adds the giving element into the current elements children tree
func (e *Markup) Apply(em *Markup) {
	if em == nil {
		return
	}

	em.AddChild(e)
}

//==============================================================================

// DeferredApply will apply a giving set of values to a giving root.
type DeferredApply struct {
	Request Appliable
	Adders  []Appliable
}

// ApplyTo adds the giving Appliables to the target during calls to Apply.
func ApplyTo(child Appliable, children ...Appliable) DeferredApply {
	return DeferredApply{
		Request: child,
		Adders:  children,
	}
}

// Apply will attempt to search for the target within the root and append
// to that giving child else ignoring the call.
func (d DeferredApply) Apply(em *Markup) {
	if em == nil {
		return
	}

	var dummy Markup
	dummy.allowChildren = true

	d.Request.Apply(&dummy)

	// If nothing was added simply ignore.
	if len(dummy.children) == 0 {
		return
	}

	child := dummy.children[0]

	// Apply new kids.
	for _, kid := range d.Adders {
		kid.Apply(child)
	}

	// Apply to the real desired root.
	child.Apply(em)
}

//==============================================================================

// AppliableTarget defines a struct which takes a giving appliable and target
// attempting to add the Appliable to the roots target child.
type AppliableTarget struct {
	Request  Appliable
	Target   string
	Multiple bool
}

// ApplyIn adds the giving Appliable to the giving target/targets if multiple is set to true
// and many are found and returns the a new Appliable which will perform the action.
func ApplyIn(child Appliable, target string, multiple bool) AppliableTarget {
	return AppliableTarget{
		Target:   target,
		Request:  child,
		Multiple: multiple,
	}
}

// Apply will attempt to search for the target within the root and append
// to that giving child else ignoring the call.
func (e AppliableTarget) Apply(em *Markup) {
	if em == nil {
		return
	}

	switch e.Multiple {
	case true:
		for _, child := range Query.QueryAll(em, e.Target) {
			e.Request.Apply(child)
		}
		break
	case false:
		if target := Query.Query(em, e.Target); target != nil {
			e.Request.Apply(target)
		}
	}
}

//==============================================================================

// Clonable defines an interface for objects that can be cloned
type Clonable interface {
	CopyTo(*Markup)
	Clone() *Markup
}

// CopyTo transfers all attributes to a giving root from the called base.
func (e *Markup) CopyTo(co *Markup) {

	//copy over the textContent
	// TODO: Should we not check if we should swap textcontent?
	// if co.textContent == "" {
	co.textContent = e.textContent
	co.textContentFn = e.textContentFn
	// }

	//clone the internal styles
	if co.allowStyles {
		for _, so := range e.styles {
			so.Clone().Apply(co)
		}
	}

	//clone the internal attribute
	if co.allowAttributes {
		for _, ao := range e.attrs {
			if name, _ := ao.Render(); name == "data-gen" {
				continue
			}

			ao.Clone().Apply(co)
		}
	}

	// co.allowAttributes = e.allowAttributes
	//clone the internal children
	if co.allowChildren {
		for _, ch := range e.children {
			ch.Clone().Apply(co)
		}
	}

	if co.allowEvents {
		for _, ch := range e.events {
			ch.Clone().Apply(co)
		}
	}

	co.morphers = append(co.morphers, e.morphers...)
}

// Clone makes a new copy of the markup structure
func (e *Markup) Clone() *Markup {
	co := NewMarkup(e.Name(), e.AutoClosed())

	//copy over the textContent
	co.textContent = e.textContent
	co.textContentFn = e.textContentFn
	co.ID = e.ID
	co.hash = e.hash
	co.uid = e.uid

	//copy over the attribute lockers
	co.allowChildren = e.allowChildren
	co.allowEvents = e.allowEvents
	co.allowAttributes = e.allowAttributes

	if e.Removed() {
		co.Removed()
	}

	//clone the internal styles
	for _, so := range e.styles {
		so.Clone().Apply(co)
	}

	co.allowStyles = e.allowStyles

	//clone the internal attribute
	for _, ao := range e.attrs {
		if name, _ := ao.Render(); name == "data-gen" {
			continue
		}

		ao.Clone().Apply(co)
	}

	// co.allowAttributes = e.allowAttributes
	//clone the internal children
	for _, ch := range e.children {
		ch.Clone().Apply(co)
	}

	for _, ch := range e.events {
		ch.Clone().Apply(co)
	}

	co.morphers = append(co.morphers, e.morphers...)

	return co
}
