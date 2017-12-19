// Package style is generated to contain the content of a css template for generating a styleguide for use in projects.

// Document is auto-generate and should not be modified by hand.

//go:generate go run generate.go

package styleguide

// styleTemplate contains the text template used to generated the full set of 
// css template for a giving styleguide.
var styleTemplate = `
html {
	font-size: {{ .BaseFontSize }}px;
	font-family: "Noto", "Roboto", Helvetica, sans-serif, serif;
}

/*
____________ Base  classes ____________________________
 Base classes for the styleguide project.
*/

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

.item-center {
    display: flex;
    flex-direction: row;
    align-self:center;
}

.content-center {
    display: flex;
    flex-direction: row;
    justify-content: center;
    align-self:center;
    align-items: center;
}

h1, h2, h3, h4, h5, h6, p {
	margin: 0;
	padding: 0;
}

a, a:hover, a:active, a:visited{
	color: #000;
	text-decoration: none;
}


/*
____________ Line height  classes ____________________________
	Line height defines giving sets of the first capital length hight
	suggestable for styles

	In typography, cap height refers to the height of a capital letter above the baseline for a particular typeface.

	print = 1.25em, Desktop(Large, Normal) = 1.3756, Tablet(Large) = 1.3756, Phone = 1.25

*/

.line-height-print, .line-height-phone {
	line-height: 1.25;
}

.line-height-deskop, .line-height-desktop-large {
	line-height: 1.3756;
}

.line-height-tablet {
	line-height: 1.3756;
}

/*
____________ Cap height  classes ____________________________
	Cap height defines giving sets of the first capital length hight
	suggestable for styles

	In typography, cap height refers to the height of a capital letter above the baseline for a particular typeface.

	small = 0.8, medium = 0.8, large = 0.68, extra-large = 0.68
*/



/*
____________ Line height  classes ____________________________
 Letter spacing class go from highest to lowest where the highest is in positive
   spacing value and the last in negative spacing values.
*/

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

/*
____________ Base border radius classes ____________________________

  These are classes for different border radius effect chosen specifically for
  use with different components.

  .smallRadius: For basic border radius for small elements eg radio, checkboxes.
  .mediumRadius: For components like text, input, labels.
  .largeRadius: For components like cards, modal boxes, etc

*/

.border-radius-sm {
	-moz-border-radius: {{ .SmallBorderRadius }}px;
	-webkit-border-radius: {{ .SmallBorderRadius }}px;
	-o-border-radius: {{ .SmallBorderRadius }}px;
	border-radius: {{ .SmallBorderRadius }}px;
}

.border-radius-md {
	-moz-border-radius: {{ .MediumBorderRadius }}px;
	-webkit-border-radius: {{ .MediumBorderRadius }}px;
	-o-border-radius: {{ .MediumBorderRadius }}px;
	border-radius: {{ .MediumBorderRadius }}px;
}

.border-radius-lg {
	-moz-border-radius: {{ .LargeBorderRadius }}px;
	-webkit-border-radius: {{ .LargeBorderRadius }}px;
	-o-border-radius: {{ .LargeBorderRadius }}px;
	border-radius: {{ .LargeBorderRadius }}px;
}

.border-radius-cirle {
	-moz-border-radius: 50%;
	-webkit-border-radius: 50%;
	-o-border-radius: 50%;
	border-radius: 50%;
}

/*
____________ Base shadowdrop classes ____________________________

  These are classes for different shadow effect chosen specifically for
  use with different components.

  .shadow: For basic shadows for normal elements.
  .shadow__drop: For shadow effects for dropdown/popovers type elements.
  .shadow__hover: For shadow effects for hovers.
  .shadow__elevanted: For shadow effects for elevated modals, cards, etc

  shadow-key-umbra-opacity: 0.2 !default;
  shadow-key-penumbra-opacity: 0.14 !default;
  shadow-ambient-shadow-opacity: 0.12 !default;
*/

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
	box-shadow: {{ .BaseShadow }};
}

.shadow-dropdown {
	box-shadow: {{ .DropShadow }};
}

.shadow-hover {
	box-shadow: {{ .HoverShadow }};
}

.shadow-elevated {
	box-shadow: {{ .FloatingShadow }};
}


/* Google's material pallete color sets */

{{ range $key, $set := .MaterialPalettes }}
  {{ range $index, $color := $set }}

  {{ $count := multiply $index 100}}
.colors-background-{{$key}}-{{ add $count 100 }} {
  background: rgba({{ $color}}, 1);
}

.colors-border-{{$key}}-{{  add $count 100 }} {
  border-color: rgba({{ $color}}, 1);
}

.colors-color-{{$key}}-{{  add $count 100 }} {
  color: rgba({{ $color}}, 1);
}

  {{ end }}
{{ end }}

/*
____________ Base font size classes ____________________________

  These are classes provide a simple set of font-scale font-size
  which allow you to use for scaling based on an initial font-size
  set on a parent, they should scale well.

  font-size-sm: Defines font size for reducing sizes
  font-size-bg: Defines font size for increasing sizes using a scale eg MajorThirds.

*/
{{ range $key, $item := .SmallHeaderScale }}
.small-heading-{{add $key 1}}{{if greaterThanEqual $key 0 }}, h{{add $key 7}} {{end}}{
	font-size: {{$item}}em;
}
{{ end }}

{{ range $key, $item := .BigHeaderScale }}
.big-heading-{{add $key 1}}{{if lessThan $key 6 }}, h{{subtract 6 $key}} {{end}}{
	font-size: {{$item}}em;
}
{{ end }}

{{ range $key, $item := .SmallFontScale }}
.small-font-{{ add $key 1 }} {
	font-size: {{$item}}em;
}
{{ end }}

{{ range $key, $item := .BigFontScale }}
.big-font-{{ add $key 1 }} {
	font-size: {{$item}}em;
}
{{ end }}

/*
____________ Font scale set ____________________________
 Taken: http://typecast.com/blog/a-more-modern-scale-for-web-typography

*/

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

/*
____________ Color set ____________________________

  These are classes provide a color set based on specific brand colors
  provided, these allows us to easily generate color, background and border-color
  classes suited to provide a consistent color design for use in project.

  These brand colors are divided into:

  primary: Main color for the giving project's brand
  secondary: Secondary brand color for project.
  success: Color for successful operation or messages.
  failure: Color for failed operations or messages.


  These are further subdivided into these diffent tones/shades:

  base: The original color without any modification.

  Other color tones are graded from 10...nth where n is a multiple of 10 * index.
  The lowest grade of 10 is where the color is close to it's darkest version while
  the highest means a continous increase in luminousity.

  We further generate classes for Color, Border-Color and Background based on the division and
  subdivision.

*/

.brand-color-primary {
	color: {{.Brand.PrimaryBrand.Base}};
}

.brand-border-color-primary {
	border-color: {{.Brand.PrimaryBrand.Base}};
}

.brand-background-color-primary {
	background: {{.Brand.PrimaryBrand.Base}};
}

{{ range $index, $item := .Brand.PrimaryBrand.Grades }}
{{ $rn := add $index 1 }}
.brand-background-color-primary-{{ multiply $rn 10}} {
	background: {{$item}};
}

.brand-border-color-primary-{{ multiply $rn 10}} {
	border-color: {{$item}};
}

.brand-primary-{{ multiply $rn 10}} {
	color: {{$item}};
}
{{ end }}


.color-primary {
	color: {{.Brand.Primary.Base}};
}

.border-color-primary {
	border-color: {{.Brand.Primary.Base}};
}

.background-color-primary {
	background: {{.Brand.Primary.Base}};
}

{{ range $index, $item := .Brand.Primary.Grades }}
{{ $rn := add $index 1 }}
.background-color-primary-{{ multiply $rn 10}} {
	background: {{$item}};
}

.color-primary-{{ multiply $rn 10}} {
	color: {{$item}};
}

.border-color-primary-{{ multiply $rn 10}} {
	border-color: {{$item}};
}
{{ end }}

/*____________ Secondary color set ____________________________

*/

.color-secondary {
	color: {{.Brand.Secondary.Base}};
}

.background-color-secondary {
	background: {{.Brand.Secondary.Base}};
}

.border-color-secondary {
	background: {{.Brand.Secondary.Base}};
}

{{ range $index, $item := .Brand.Secondary.Grades }}
{{ $rn := add $index 1 }}
.background-color-secondary-{{ multiply $rn 10}} {
	background: {{$item}};
}

.border-color-secondary-{{ multiply $rn 10}} {
	border-color: {{$item}};
}

.color-secondary-{{ multiply $rn 10}} {
	color: {{$item}};
}
{{ end }}

.brand-color-secondary {
	color: {{.Brand.SecondaryBrand.Base}};
}

.brand-background-color-secondary {
	background: {{.Brand.SecondaryBrand.Base}};
}

.brand-border-color-secondary {
	background: {{.Brand.SecondaryBrand.Base}};
}

{{ range $index, $item := .Brand.SecondaryBrand.Grades }}
{{ $rn := add $index 1 }}
.brand-background-color-secondary-{{ multiply $rn 10}} {
	background: {{$item}};
}

.brand-border-color-secondary-{{ multiply $rn 10}} {
	border-color: {{$item}};
}

.brand-color-secondary-{{ multiply $rn 10}} {
	color: {{$item}};
}
{{ end }}


/*____________ Success color set ____________________________

*/

.brand-success {
	color: {{.Brand.Success.Base}};
}

.background-color-success {
	background: {{.Brand.Success.Base}};
}

.border-color-success {
	border-color: {{.Brand.Success.Base}};
}

{{ range $index, $item := .Brand.Success.Grades }}
{{ $rn := add $index 1 }}
.background-color-success-{{ multiply $rn 10}} {
	background: {{$item}};
}

.brand-success-{{ multiply $rn 10}} {
	color: {{$item}};
}

.border-color-success-{{ multiply $rn 10}} {
	border-color: {{$item}};
}
{{ end }}

/*____________ White color set ____________________________

*/

.background-color-white {
	background: {{.Brand.White.Base}};
}

.brand-white {
	color: {{.Brand.White.Base}};
}

.border-color-white {
	border-color: {{.Brand.White.Base}};
}

{{ range $index, $item := .Brand.White.Grades }}
{{ $rn := add $index 1 }}
.background-color-white-{{ multiply $rn 10}} {
	background: {{$item}};
}

.brand-white-{{ multiply $rn 10}} {
	color: {{$item}};
}

.border-color-white-{{ multiply $rn 10}} {
	border-color: {{$item}};
}
{{ end }}

/*____________ Failure color set ____________________________

*/

.background-color-failure {
	background: {{.Brand.Failure.Base}};
}

.brand-failure {
	color: {{.Brand.Failure.Base}};
}

.border-color-failure {
	border-color: {{.Brand.Failure.Base}};
}

{{ range $index, $item := .Brand.Failure.Grades }}
{{ $rn := add $index 1 }}
.background-color-failure-{{ multiply $rn 10}} {
	background: {{$item}};
}

.brand-failure-{{ multiply $rn 10}} {
	color: {{$item}};
}

.border-color-failure-{{ multiply $rn 10}} {
	border-color: {{$item}};
}
{{ end }}

/*______________________________________________________________________

*/
`
