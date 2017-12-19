package trees

var (
	// Hide will set any markup top level root into a style of display none.
	Hide hideMarkup

	// Show will set any markup top level root into a style of display block.
	Show showMarkup
)

// hide*Markup provides a markup property modifier for  setting style property
// to hidden.
type hideMarkup struct{}

// Mod marks the given markup as display:none
func (v hideMarkup) Mode(m *Markup) {
	ReplaceORAddStyle(m, "display", "none")
}

// show*Markup provides a markup property modifier for  setting style property
// to displayed.
type showMarkup struct{}

// Mod marks the given markup with a display: block
func (v showMarkup) Mode(m *Markup) {
	ReplaceORAddStyle(m, "display", "block")
}
