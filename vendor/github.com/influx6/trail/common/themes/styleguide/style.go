package styleguide

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
	"text/template"

	"github.com/influx6/trail/common"
	"github.com/influx6/faux/colors"
	colorful "github.com/lucasb-eyer/go-colorful"
)

// contains different constants used within the package.
const (
	shadowLarge    = "0px 20px 40px 4px rgba(0, 0, 0, 0.51)"
	shadowPopDrops = "0px 9px 30px 2px rgba(0, 0, 0, 0.51)"
	shadowHovers   = "0px 13px 30px 5px rgba(0, 0, 0, 0.58)"
	shadowNormal   = "0px 13px 20px 2px rgba(0, 0, 0, 0.45)"

	AnimationCurveFastOutSlowIn   = "cubic-bezier(0.4, 0, 0.2, 1)"
	AnimationCurveLinearOutSlowIn = "cubic-bezier(0, 0, 0.2, 1)"
	AnimationCurveFastOutLinearIn = "cubic-bezier(0.4, 0, 1, 1)"
	AnimationCurveDefault         = AnimationCurveFastOutSlowIn

	smallBorderRadius  = 2
	mediumBorderRadius = 4
	largeBorderRadius  = 8

	AugmentedFourth = 1.414
	MinorSecond     = 1.067
	MajorSecond     = 1.125
	MinorThird      = 1.200
	MajorThird      = 1.250
	PerfectFourth   = 1.333
	PerfectFifth    = 1.500
	GoldenRatio     = 1.618

	LuminFlat     = 1.015
	LuminFat      = 1.200
	LuminFatThird = 1.245
)

var (
	helpers = template.FuncMap{
		"quote": func(b interface{}) string {
			switch bo := b.(type) {
			case string:
				return strconv.Quote(bo)
			case int:
				return strconv.Quote(strconv.Itoa(bo))
			case int64:
				return strconv.Quote(strconv.Itoa(int(bo)))
			case float32:
				mo := strconv.FormatFloat(float64(bo), 'f', 4, 32)
				return strconv.Quote(mo)
			case float64:
				mo := strconv.FormatFloat(bo, 'f', 4, 32)
				return strconv.Quote(mo)
			case rune:
				return strconv.QuoteRune(bo)
			default:
				mo, err := json.Marshal(b)
				if err != nil {
					return err.Error()
				}

				return strconv.Quote(string(mo))
			}
		},
		"prefixInt": func(prefix string, b int) string {
			return fmt.Sprintf("%s%d", prefix, b)
		},
		"add": func(a, b int) int {
			return a + b
		},
		"lessThanEqual": func(a, b int) bool {
			return a <= b
		},
		"greaterThanEqual": func(a, b int) bool {
			return a >= b
		},
		"lessThan": func(a, b int) bool {
			return a < b
		},
		"greaterThan": func(a, b int) bool {
			return a > b
		},
		"len": func(a interface{}) int {
			switch real := a.(type) {
			case []interface{}:
				return len(real)
			case [][]byte:
				return len(real)
			case []byte:
				return len(real)
			case []float32:
				return len(real)
			case []float64:
				return len(real)
			case []string:
				return len(real)
			case []int:
				return len(real)
			default:
				return 0
			}
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
		"textRhythmn": func(lineHeight int, capHeight int, fontSize int) int {
			return ((lineHeight - capHeight) * fontSize) / 2
		},
		"textRhythmnEM": func(lineHeight, capHeight, fontSize float64) float64 {
			return ((lineHeight - capHeight) * fontSize) / 2
		},
	}
)

// StyleColors defines a struct which holds all possible giving brand colors utilized
// for the project.
type styleColors struct {
	Primary        Tones `json:"primary"`
	Secondary      Tones `json:"secondary"`
	Success        Tones `json:"success"`
	Failure        Tones `json:"failure"`
	White          Tones `json:"white"`
	PrimaryBrand   Tones `json:"primary_support"`
	SecondaryBrand Tones `json:"secondary_support"`
}

// Render initializes the style guide and all internal properties into
// appropriate defaults and states and generates a css style written into
// the provided writer
func Render(w io.Writer, attr common.Theme) error {
	var err error

	var brand styleColors

	attr = initAttr(attr)

	if attr.PrimaryBrandColor != "" {
		brand.PrimaryBrand, err = NewTones(attr.PrimaryBrandColor)
		if err != nil {
			return errors.New("Invalid primary brand color: " + err.Error())
		}
	}

	if attr.SecondaryBrandColor != "" {
		brand.SecondaryBrand, err = NewTones(attr.SecondaryBrandColor)
		if err != nil {
			return errors.New("Invalid secondary brand color: " + err.Error())
		}
	}

	brand.Primary, err = NewTones(attr.PrimaryColor)
	if err != nil {
		return errors.New("Invalid primary color: " + err.Error())
	}

	brand.Secondary, err = NewTones(attr.SecondaryColor)
	if err != nil {
		return errors.New("Invalid secondary color: " + err.Error())
	}

	brand.White, err = NewTones(attr.PrimaryWhite)
	if err != nil {
		return errors.New("Invalid white color: " + err.Error())
	}

	brand.Success, err = NewTones(attr.SuccessColor)
	if err != nil {
		return errors.New("Invalid success color: " + err.Error())
	}

	brand.Failure, err = NewTones(attr.FailureColor)
	if err != nil {
		return errors.New("Invalid failure color: " + err.Error())
	}

	tml, err := template.New("styleguide").Funcs(helpers).Parse(styleTemplate)
	if err != nil {
		return err
	}

	shm, bhm := GenerateValueScale(1, attr.HeaderBaseScale, attr.MinimumHeadScaleCount, attr.MaximumHeadScaleCount)
	sm, bg := GenerateValueScale(1, attr.BaseScale, attr.MinimumScaleCount, attr.MaximumScaleCount)

	return tml.Execute(w, struct {
		common.Theme
		Brand            styleColors
		SmallFontScale   []float64
		BigFontScale     []float64
		SmallHeaderScale []float64
		BigHeaderScale   []float64
	}{
		SmallFontScale:   sm,
		BigFontScale:     bg,
		SmallHeaderScale: shm,
		BigHeaderScale:   bhm,
		Brand:            brand,
		Theme:            attr,
	})
}

//================================================================================================

// Tones defines the set of color tones generated for a base color using the Hamonic tone
// sets, it provides a very easily set of color variations for use in styles.
type Tones struct {
	Base   Color   `json:"base"`
	Grades []Color `json:"tones"`
}

// NewTones returns a new Tones object representing the provided color tones if
// the value provided is a valid color.
func NewTones(base string) (Tones, error) {
	c, err := ColorFrom(base)
	if err != nil {
		return Tones{}, err
	}

	return HamonicsFrom(c), nil
}

// String returns the string representation of the provided tone.
func (t Tones) String() string {
	return fmt.Sprintf(`%q %q`, t.Base, t.Grades)
}

func initAttr(attr common.Theme) common.Theme {
	if attr.MaterialPalettes == nil || len(attr.MaterialPalettes) == 0 {
		attr.MaterialPalettes = MaterialPalettes
	}

	if attr.PrimaryColor == "" {
		attr.PrimaryColor = fmt.Sprintf("rgb(%s)", MaterialPalettes["blue"][5])
	}

	if attr.SecondaryColor == "" {
		attr.SecondaryColor = fmt.Sprintf("rgb(%s)", MaterialPalettes["deep-purple"][5])
	}

	if attr.PrimaryWhite == "" {
		attr.PrimaryWhite = fmt.Sprintf("rgb(%s)", MaterialPalettes["white"][0])
	}

	if attr.SuccessColor == "" {
		attr.SuccessColor = fmt.Sprintf("rgb(%s)", MaterialPalettes["green"][5])
	}

	if attr.FailureColor == "" {
		attr.FailureColor = fmt.Sprintf("rgb(%s)", MaterialPalettes["red"][5])
	}

	if attr.BaseFontSize <= 0 {
		attr.BaseFontSize = 16
	}

	if attr.BaseScale <= 0 {
		attr.BaseScale = PerfectFourth
	}

	if attr.HeaderBaseScale <= 0 {
		attr.HeaderBaseScale = MajorThird
	}

	if attr.MinimumHeadScaleCount == 0 {
		attr.MinimumHeadScaleCount = 4
	}

	if attr.MaximumHeadScaleCount == 0 {
		attr.MaximumHeadScaleCount = 6
	}

	if attr.MinimumScaleCount == 0 {
		attr.MinimumScaleCount = 10
	}

	if attr.MaximumScaleCount == 0 {
		attr.MaximumScaleCount = 10
	}

	if attr.AnimationCurveDefault == "" {
		attr.AnimationCurveDefault = AnimationCurveDefault
	}

	if attr.AnimationCurveFastOutLinearIn == "" {
		attr.AnimationCurveFastOutLinearIn = AnimationCurveFastOutLinearIn
	}

	if attr.AnimationCurveFastOutSlowIn == "" {
		attr.AnimationCurveFastOutSlowIn = AnimationCurveFastOutSlowIn
	}

	if attr.AnimationCurveLinearOutSlowIn == "" {
		attr.AnimationCurveLinearOutSlowIn = AnimationCurveLinearOutSlowIn
	}

	if attr.FloatingShadow == "" {
		attr.FloatingShadow = shadowLarge
	}

	if attr.HoverShadow == "" {
		attr.HoverShadow = shadowHovers
	}

	if attr.BaseShadow == "" {
		attr.BaseShadow = shadowNormal
	}

	if attr.DropShadow == "" {
		attr.DropShadow = shadowPopDrops
	}

	if attr.SmallBorderRadius <= 0 {
		attr.SmallBorderRadius = smallBorderRadius
	}

	if attr.MediumBorderRadius <= 0 {
		attr.MediumBorderRadius = mediumBorderRadius
	}

	if attr.LargeBorderRadius <= 0 {
		attr.LargeBorderRadius = largeBorderRadius
	}

	return attr
}

//================================================================================================

// HamonicsFrom uses the above scale to return a slice of new Colors based on the provided
// HamonyScale set.
func HamonicsFrom(c Color) Tones {

	var scale []float64

	min, max := GenerateValueScale(0.1, LuminFatThird, 1, 10)

	lastItem := max[len(max)-1]
	_, inmax := GenerateValueScale(lastItem, LuminFlat, 0, 8)

	inmax = inmax[1:]

	reverse(len(min), func(index int) {
		scale = append(scale, min[index])
	})

	scale = append(scale, max...)
	scale = append(scale, inmax...)

	var colors []Color

	// TODO(alex): Should we have another scale for saturation?
	for _, scale := range scale {
		if scale > 1 {
			scale = 1
		}

		newColor := colorful.Hsl(c.Hue, c.Saturation, scale)
		h, s, l := newColor.Hsl()

		colors = append(colors, Color{
			C:          newColor,
			Hue:        h,
			Saturation: s,
			Luminosity: l,
			Alpha:      c.Alpha,
		})
	}

	var t Tones
	t.Base = c
	t.Grades = colors

	return t
}

// AdditiveSaturation adds the provided scale to the colors saturation value
// returning a new color suited to match.
func AdditiveSaturation(c Color, scale float64) Color {
	newLumen := c.Saturation + scale

	if newLumen > 1 {
		newLumen = 1
	}

	if newLumen < 0 {
		newLumen = 0
	}

	newColor := colorful.Hsl(c.Hue, newLumen, c.Luminosity)

	h, s, l := newColor.Hsl()

	return Color{
		C:          newColor,
		Hue:        h,
		Saturation: s,
		Luminosity: l,
		Alpha:      c.Alpha,
	}
}

// MultiplicativeSaturation multiples the scale to the colors saturation value
// using the returned value as a addition to the current saturation value,
// Creating a gradual change in saturation for the returned color.
func MultiplicativeSaturation(c Color, scale float64) Color {
	newLuma := (c.Saturation * scale)
	newLumen := c.Saturation + newLuma

	if newLumen > 1 {
		newLumen = 1
	}

	if newLumen < 0 {
		newLumen = 0
	}

	newColor := colorful.Hsl(c.Hue, newLumen, c.Luminosity)
	h, s, l := newColor.Hsl()

	return Color{
		C:          newColor,
		Hue:        h,
		Saturation: s,
		Luminosity: l,
		Alpha:      c.Alpha,
	}
}

// AdditiveLumination adds the provided scale to the colors Luminouse value
// returning a new color suited to match.
func AdditiveLumination(c Color, scale float64) Color {
	newLumen := c.Luminosity + scale

	if newLumen > 1 {
		newLumen = 1
	}

	if newLumen < 0 {
		newLumen = 0
	}

	newColor := colorful.Hsl(c.Hue, c.Saturation, newLumen)

	h, s, l := newColor.Hsl()

	return Color{
		C:          newColor,
		Hue:        h,
		Saturation: s,
		Luminosity: l,
		Alpha:      c.Alpha,
	}
}

// MultiplicativeLumination multiples the scale to the colors Luminouse value
// using the returned value as addition to the current Luminouse value.
// Creating a gradual change in luminousity for the returned color.
func MultiplicativeLumination(c Color, scale float64) Color {
	// fmt.Printf("ML: H: %.4f S: %.4f L: %.4f \n", c.Hue, c.Saturation, c.Luminosity)
	newLum := (c.Luminosity * scale)
	newLumen := c.Luminosity + newLum

	if newLumen > 1 {
		newLumen = 1
	}

	if newLumen < 0 {
		newLumen = 0
	}

	newColor := colorful.Hsl(c.Hue, c.Saturation, newLumen)
	h, s, l := newColor.Hsl()

	return Color{
		C:          newColor,
		Hue:        h,
		Saturation: s,
		Luminosity: l,
		Alpha:      c.Alpha,
	}
}

//================================================================================================

// Color defines a basic struct which expresses the color values provided
// a struct containing HSL points.
type Color struct {
	C          colorful.Color
	Hue        float64 `json:"hue"`
	Luminosity float64 `json:"luminosity"`
	Saturation float64 `json:"saturation"`
	Alpha      float64 `json:"alpha"`
}

// String returns the Hex representation of the color.
func (c Color) String() string {
	return c.C.Hex()
}

// ColorFrom returns a Color instance representing the valid
// color values provided else returning error if the color value
// is not a valid color presentation i.e (rgb,rgba, hsl, hex).
func ColorFrom(value string) (Color, error) {
	var c colorful.Color

	alpha := float64(1)

	switch {
	case colors.IsHex(value):
		c, _ = colorful.Hex(value)
		break
	case colors.IsHSL(value):
		h, s, l := colors.ParseHSL(value)
		c = colorful.Hsl(h, s, l)
		break
	case colors.IsRGB(value):
		var red, green, blue int
		red, green, blue, alpha = colors.ParseRGB(value)
		c = colorful.Color{R: float64(red) / 255, G: float64(green) / 255, B: float64(blue) / 255}
		break
	case colors.IsRGBA(value):
		var red, green, blue int
		red, green, blue, alpha = colors.ParseRGB(value)
		c = colorful.Color{R: float64(red) / 255, G: float64(green) / 255, B: float64(blue) / 255}
		break
	default:
		return Color{}, errors.New("Invalid color value received")
	}

	h, s, l := c.Hsl()

	return Color{
		C:          c,
		Hue:        h,
		Saturation: s,
		Luminosity: l,
		Alpha:      alpha,
	}, nil
}

//==================================================================================

// GenerateValueScale returns a slice of values which are the a combination of
// a reducing + increasing scaled values of the provided scale generated from
// using the base initial 1.0 value against an ever incremental 1.0*(scale * n)
// or 1.0 / (scale *n) value, where n is the ever increasing index.
func GenerateValueScale(base float64, scale float64, minorCount int, majorCount int) ([]float64, []float64) {
	var major, minor []float64

	times(minorCount, func(index int) {
		if index > 1 {
			prevValue := minor[len(minor)-1]
			minor = append(minor, prevValue/scale)
			return
		}

		minor = append(minor, base/scale)
	})

	major = append(major, base)

	times(majorCount, func(index int) {
		if index > 1 {
			prevValue := major[index-1]
			major = append(major, prevValue*scale)
			return
		}

		major = append(major, base*scale)
	})

	return minor, major
}

func times(n int, fn func(int)) {
	for i := 0; i < n; i++ {
		fn(i + 1)
	}
}

func reverse(n int, fn func(int)) {
	for i := n; i > 0; i-- {
		fn(i - 1)
	}
}
