package trees

import (
	"bytes"
	"crypto/rand"
	"strings"
	"text/template"
)

var (
	helpers = template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
		"multiply": func(a, b int) int {
			return a * b
		},
		"subtract": func(a, b int) int {
			return a - b
		},
		"divide": func(a, b int) int {
			return a / b
		},
		"perc": func(a, b float64) float64 {
			return (a / b) * 100
		},
	}
)

// Templated takes a giving string, parses into a template, runs against the processor and
// parses the provided result into a markup.
func Templated(tml string, bind interface{}, processor func(string) string) (string, error) {
	tmp, err := template.New("templated").Funcs(helpers).Parse(tml)
	if err != nil {
		return "", err
	}

	var content bytes.Buffer
	if err := tmp.Execute(&content, bind); err != nil {
		return "", err
	}

	return processor(content.String()), nil
}

// RandString generates a set of random numbers of a set length
func RandString(n int) string {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, n)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}

// Augment adds new markup to an the root if its Element
func Augment(root *Markup, m ...*Markup) {
	for _, mo := range m {
		mo.Apply(root)
	}
}

// Events defines an interface that returns a list of events.
type Events interface {
	Events() []Event
}

// Styles defines an interface that returns a list of styles.
type Styles interface {
	Styles() []Property
}

// EqualStyles returns true/false if the style values are all equal attribute.
func EqualStyles(e, em Styles) bool {
	old := em.Styles()

	if len(old) < 1 {
		if len(e.Styles()) > 0 {
			return false
		}
		return true
	}

	for _, oa := range old {
		name, val := oa.Render()

		ta, err := GetStyle(e, name)
		if err != nil {
			return false
		}

		_, tvalue := ta.Render()
		if tvalue != val {
			return false
		}
	}

	return true
}

// Attributes defines an interface that returns a list of attributes.
type Attributes interface {
	Attributes() []Property
}

// EqualAttributes returns true/false if the elements and the giving markup have equal attribute
func EqualAttributes(e, em Attributes) bool {
	old := em.Attributes()

	if len(old) < 1 {
		if len(e.Attributes()) > 0 {
			return false
		}
		return true
	}

	for _, oa := range old {
		name, val := oa.Render()

		ta, err := GetAttr(e, name)
		if err != nil {
			return false
		}

		_, tvalue := ta.Render()
		if tvalue != val {
			return false
		}
	}

	return true
}

// GetStyles returns the styles that contain the specified name and if not empty that contains the specified value also, note that strings
// NOTE: string.Contains is used when checking value parameter if present
func GetStyles(e Styles, f, val string) []Property {
	var found []Property
	var styles = e.Styles()

	for _, as := range styles {
		name, value := as.Render()
		if name != f {
			continue
		}

		if val != "" {
			if !strings.Contains(value, val) {
				continue
			}
		}

		found = append(found, as)
	}

	return found
}

// GetStyle returns the style with the specified tag name
func GetStyle(e Styles, f string) (Property, error) {
	styles := e.Styles()
	for _, as := range styles {
		name, _ := as.Render()
		if name == f {
			return as, nil
		}
	}
	return nil, ErrNotFound
}

// StyleContains returns the styles that contain the specified name and if the val is not empty then
// that contains the specified value also, note that strings
// NOTE: string.Contains is used
func StyleContains(e Styles, f, val string) bool {
	styles := e.Styles()
	for _, as := range styles {
		name, value := as.Render()
		if !strings.Contains(name, f) {
			continue
		}

		if val != "" {
			if !strings.Contains(value, val) {
				continue
			}
		}

		return true
	}

	return false
}

// GetAttrs returns the attributes that have the specified text within the naming
// convention and if it also contains the set val if not an empty "",
// NOTE: string.Contains is used
func GetAttrs(e Attributes, f, val string) []Property {
	var found []Property

	for _, as := range e.Attributes() {
		name, value := as.Render()
		if name != f {
			continue
		}

		if val != "" {
			if !strings.Contains(value, val) {
				continue
			}
		}

		found = append(found, as)
	}

	return found
}

// AttrContains returns the attributes that have the specified text within the naming
// convention and if it also contains the set val if not an empty "",
// NOTE: string.Contains is used
func AttrContains(e Attributes, f, val string) bool {
	for _, as := range e.Attributes() {
		name, value := as.Render()
		if !strings.Contains(name, f) {
			continue
		}

		if val != "" {
			if !strings.Contains(value, val) {
				continue
			}
		}

		return true
	}

	return false
}

// GetAttr returns the attribute with the specified tag name
func GetAttr(e Attributes, f string) (Property, error) {
	for _, as := range e.Attributes() {
		name, _ := as.Render()
		if name == f {
			return as, nil
		}
	}
	return nil, ErrNotFound
}

// ReplaceStyle replaces a specific style with the given
// name with the supplied value.
func ReplaceStyle(m Styles, name string, val string) {
	styl, err := GetStyle(m, name)
	if err != nil {
		return
	}

	stylm, ok := styl.(*CSSStyle)
	if !ok {
		return
	}

	stylm.Value = val
}

// ReplaceAttribute replaces a specific attribute with the given
// name with the supplied value.
func ReplaceAttribute(m Attributes, name string, val string) {
	attr, err := GetAttr(m, name)
	if err != nil {
		return
	}

	attrm, ok := attr.(*Attribute)
	if !ok {
		return
	}

	attrm.Value = val
}

// ReplaceORAddStyle replaces a specific style with the given
// name with the supplied value if not found it adds a new one
// if found and if the type does not match a *CSSStyle then it stops.
func ReplaceORAddStyle(m Properties, name string, val string) {
	styl, err := GetStyle(m, name)
	if err != nil {
		m.AddStyle(NewCSSStyle(name, val))
		return
	}

	stylm, ok := styl.(*CSSStyle)
	if !ok {
		return
	}

	stylm.Value = val
}

// ReplaceORAddAttribute replaces a specific attribute with the given
// name with the supplied value if not found it adds a new one
// if found and if the type does not match a *CSSStyle then it stops.
func ReplaceORAddAttribute(m Properties, name string, val string) {
	attr, err := GetAttr(m, name)
	if err != nil {
		m.AddAttribute(NewAttr(name, val))
		return
	}

	if attrm, ok := attr.(*Attribute); ok {
		attrm.Value = val
		return
	}

	if classlist, ok := attr.(*ClassList); ok {
		classlist.list = nil
		classlist.list = append(classlist.list, val)
	}
}

//==============================================================================

// ElementsUsingStyle returns the children within the element matching the
// stlye restrictions passed.
// NOTE: is uses StyleContains
func ElementsUsingStyle(root *Markup, key string, val string) []*Markup {
	var found []*Markup

	for _, ch := range root.Children() {
		if StyleContains(ch, key, val) {
			found = append(found, ch)
		}

		found = append(found, ElementsUsingStyle(ch, key, val)...)
	}

	return found
}

// ElementsWithAttr returns the children within the element matching the
// stlye restrictions passed.
// NOTE: is uses AttrContains
func ElementsWithAttr(root *Markup, key string, val string) []*Markup {
	var found []*Markup

	for _, ch := range root.Children() {
		if AttrContains(ch, key, val) {
			found = append(found, ch)
		}

		found = append(found, ElementsWithAttr(ch, key, val)...)
	}

	return found
}

// ElementsWithTag returns elements matching the tag type in the parent markup children list
// only without going deeper into children's children lists.
func ElementsWithTag(root *Markup, tag string) []*Markup {
	var found []*Markup

	tag = strings.TrimSpace(strings.ToLower(tag))
	for _, ch := range root.Children() {
		if ch.Name() == tag {
			found = append(found, ch)
		}

		found = append(found, ElementsWithTag(ch, tag)...)
	}

	return found
}

//==============================================================================
