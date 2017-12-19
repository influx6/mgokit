package css

import (
	bcss "github.com/aymerick/douceur/css"
	"github.com/aymerick/douceur/parser"
)

const (
	base = `
.line-height-print, .line-height-phone {
	line-height: 1.25;
}

.line-height-deskop, .line-height-desktop-large {
	line-height: 1.3756;
}

.line-height-tablet {
	line-height: 1.3756;
}


.letter-spacing-1 {
	letter-spacing: 1px;
}

.letter-spacing-2 {
	letter-spacing: -0.5px;
}

.letter-spacing-3 {
	letter-spacing: -1px;
}

.letter-spacing-4 {
	letter-spacing: -2px;
}

.focus-shadow {
  box-shadow: 0 0 8px rgba(0,0,0,.18),0 8px 16px rgba(0,0,0,.36);
}

.shadow-2dp {
	box-shadow: 0 2px 2px 0 rgba(0, 0, 0, 0.14),
              0 3px 1px -2px rgba(0, 0, 0, 0.2),
              0 1px 5px 0 rgba(0, 0, 0, 0.12);
}

.shadow-3dp {
  box-shadow: 0 3px 4px 0 rgba(0, 0, 0, 0.14),
              0 3px 3px -2px rgba(0, 0, 0, 0.2),
              0 1px 8px 0 rgba(0, 0, 0, 0.12);
}

.shadow-4dp {
  box-shadow: 0 4px 5px 0 rgba(0, 0, 0, 0.14),
              0 1px 10px 0 rgba(0, 0, 0, 0.12),
              0 2px 4px -1px rgba(0, 0, 0, 0.2);
}

.shadow-6dp {
  box-shadow: 0 6px 10px 0 rgba(0, 0, 0, 0.14),
              0 1px 18px 0 rgba(0, 0, 0, 0.12),
              0 3px 5px -1px rgba(0, 0, 0, 0.2);
}

.shadow-8dp {
  box-shadow: 0 8px 10px 1px rgba(0, 0, 0, 0.14),
              0 3px 14px 2px rgba(0, 0, 0, 0.12),
              0 5px 5px -3px rgba(0, 0, 0, 0.2);
}

.shadow-16dp {
  box-shadow: 0 16px 24px 2px rgba(0, 0, 0, 0.14),
              0  6px 30px 5px rgba(0, 0, 0, 0.12),
              0  8px 10px -5px rgba(0, 0, 0, 0.2);
}

.shadow-24dp {
  box-shadow: 0  9px 46px  8px rgba(0, 0, 0, 0.14),
              0 11px 15px -7px rgba(0, 0, 0, 0.12),
              0 24px 38px  3px rgba(0, 0, 0, 0.2);
}

.shadow {
	box-shadow: 0px 13px 20px 2px rgba(0, 0, 0, 0.45);
}

.shadow-dropdown {
	box-shadow: 0px 9px 30px 2px rgba(0, 0, 0, 0.51);
}

.shadow-hover {
	box-shadow: 0px 13px 30px 5px rgba(0, 0, 0, 0.58);
}

.shadow-elevated {
	box-shadow: 0px 20px 40px 4px rgba(0, 0, 0, 0.51);
}

.wrap {
  text-wrap: wrap;
  white-space: -moz-pre-wrap;
  white-space: pre-wrap;
  word-wrap: break-word;
}

.sizing {
  box-sizing: border-box;
  -webkit-box-sizing: border-box;
  -o-box-sizing: border-box;
  -moz-box-sizing: border-box;
}

.clear {
	content: " ";
	clear: both;
	display: block;
	visibility: hidden;
	height: 0;
	font-size: 0;
}

.body-desktop {
	font-size: 16px;
	font-size: 1em;
	line-height: 1.375em;
}

.body-desktop-lg {
	font-size: 16px;
	font-size: 1em;
	line-height: 1.25em;
}

.body-tablet-lg {
	font-size: 16px;
	font-size: 1em;
	line-height: 1.375em;
}

.body-tablet-sm {
	font-size: 16px;
	font-size: 1em;
	line-height: 1.25em;
}

.body-phone {
	font-size: 16px;
	font-size: 1em;
	line-height: 1.25em;
}

.title-font {
  font-size: 20px;
  font-weight: 500;
  line-height: 1;
  letter-spacing: 0.02em;
}

.title-font-contrast {
    opacity: 0.87;
}

.subheadline-font {
  font-size: 16px;
  font-weight: 400;
  line-height: 24px;
  letter-spacing: 0.04em;
}

.subheadline-font-contrast {
    opacity: 0.87;
}

.subheadline2-font {
  font-size: 16px;
  font-weight: 400;
  line-height: 28px;
  letter-spacing: 0.04em;
}

.subheadline2-font-contrast {
    opacity: 0.87;
}

.body-font {
  font-size: 14px;
  line-height: 24px;
  letter-spacing: 0;
  font-weight: 500;
}

.body-font-contrast {
    opacity: 0.87;
}

.body1-font {
  font-size: 14px;
  font-weight: 400;
  line-height: 24px;
  letter-spacing: 0;
}

.body1-font-contrast {
    opacity: 0.87;
}

.headline-font {
  font-size: 24px;
  font-weight: 400;
  line-height: 32px;
  -moz-osx-font-smoothing: grayscale;
}

.headline-font-contrast {
    opacity: 0.87;
}

.caption-font {
  font-size: 12px;
  font-weight: 400;
  line-height: 1;
  letter-spacing: 0;
}

.caption-font-contrast {
    opacity: 0.54;
}

.caption-font-contrast {
    opacity: 0.54;
}

.blockquote-font {
  position: relative;
  font-size: 24px;
  font-weight: 300;
  font-style: italic;
  line-height: 1.35;
  letter-spacing: 0.08em;
}

.blockquote-font-contrast {
    opacity: 0.54;
}

.blockquote-font:before, .blockquote-font-contrast:before {
    position: absolute;
    left: -0.5em;
    content: '“';
}

.blockquote-font:after, .blockquote-font-contrast:after {
    content: '”';
    margin-left: -0.05em;
}

.menu-font {
  font-size: 14px;
  font-weight: 500;
  line-height: 1;
  letter-spacing: 0;
}

.menu-font-with-contrast {
  font-size: 14px;
  font-weight: 500;
  line-height: 1;
  letter-spacing: 0;
  opacity: 0.87;
}

.button-font {
  font-size: 14px;
  font-weight: 500;
  text-transform: uppercase;
  line-height: 1;
  letter-spacing: 0;
}

.button-font-with-contrast {
  font-size: 14px;
  font-weight: 500;
  text-transform: uppercase;
  line-height: 1;
  letter-spacing: 0;
  opacity: 0.87;
}

.material-icons {
  font-family: 'Material Icons';
  font-weight: normal;
  font-style: normal;
  font-size: 24px;
  line-height: 1;
  letter-spacing: normal;
  text-transform: none;
  display: inline-block;
  word-wrap: normal;
  font-feature-settings: 'liga';
  -webkit-font-feature-settings: 'liga';
  -webkit-font-smoothing: antialiased;
}
	
`
)

var (
	baseStyles = func() bcss.Stylesheet {
		if extend, err := parser.Parse(base); err == nil {
			return *extend
		}

		return bcss.Stylesheet{}
	}()
)
