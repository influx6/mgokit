package css

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"text/template"

	bcss "github.com/aymerick/douceur/css"
	"github.com/aymerick/douceur/parser"
)

var (
	animationCurveFastOutSlowIn   = "cubic-bezier(0.4, 0, 0.2, 1)"
	animationCurveLinearOutSlowIn = "cubic-bezier(0, 0, 0.2, 1)"
	animationCurveFastOutLinearIn = "cubic-bezier(0.4, 0, 1, 1)"
	animationCurveDefault         = animationCurveFastOutSlowIn
	helpers                       = template.FuncMap{
		"materialColors": func(colorName string, grade int) string {
			colorName = strings.ToLower(colorName)

			if grade < 0 {
				grade = 0
			}

			wantedColor, ok := materialPalettes[colorName]
			if !ok {
				return "rgba(0,0,0,1)"
			}

			var colorVals string

			if grade >= len(wantedColor) {
				colorVals = wantedColor[len(wantedColor)-1]
			} else {
				colorVals = wantedColor[grade]
			}

			return fmt.Sprintf("rgba(%s,1)", colorVals)
		},
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
		"multiplyf": func(a, b float64) float64 {
			return a * b
		},
		"subtractf": func(a, b float64) float64 {
			return a - b
		},
		"dividef": func(a, b float64) float64 {
			return a / b
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
		"dialogWidth": func(unit int) string {
			return fmt.Sprintf(`
				width: %dpx;
			`, unit*56)
		},
		"animationDefaultProperty": func() string {
			return animationCurveDefault
		},
		"animationFastOutLinearInProperty": func() string {
			return animationCurveFastOutLinearIn
		},
		"animationFastOutSlowInProperty": func() string {
			return animationCurveFastOutSlowIn
		},
		"animationLinearOutSlowInProperty": func() string {
			return animationCurveLinearOutSlowIn
		},
		"animationDefault": func(duration float32) string {
			if duration < 0 {
				duration = 0.2
			}

			return fmt.Sprintf(`
				transition-duration: %.4fs;
				transition-timing-function: %s;
			`, duration, animationCurveDefault)
		},
		"animationFastOutLinearIn": func(duration float32) string {
			if duration < 0 {
				duration = 0.2
			}

			return fmt.Sprintf(`
				transition-duration: %.4fs;
				transition-timing-function: %s;
			`, duration, animationCurveFastOutLinearIn)
		},
		"animationFastOutSlowIn": func(duration float32) string {
			if duration < 0 {
				duration = 0.2
			}

			return fmt.Sprintf(`
				transition-duration: %.4fs;
				transition-timing-function: %s;
			`, duration, animationCurveFastOutSlowIn)
		},
		"animationLinearOutSlowIn": func(duration float32) string {
			if duration < 0 {
				duration = 0.2
			}

			return fmt.Sprintf(`
				transition-duration: %.4fs;
				transition-timing-function: %s;
			`, duration, animationCurveLinearOutSlowIn)
		},
	}
)

// Rule defines the a single css rule which will be transformed and
// converted into a usable stylesheet during rendering.
type Rule struct {
	plain     string
	feed      *Rule
	depends   []*Rule
	feedStyle *bcss.Stylesheet
	template  *template.Template
}

// New returns a new instance of a Rule which provides capability to parse
// and extrapolate the giving content using the provided binding.
// - Arguments:
// 		- rules : text containing css values.
//		- extension: A instance of a Rule, that may contain certain styles which can be extended into current rule styles using the `extend` template function.
// 		- rules: A slice of rules which should be built with this, they will also inherit this rules parents, a nice way to
// 				extend a rule sets property.
func New(rules string, extension *Rule, rs ...*Rule) *Rule {
	rsc := &Rule{depends: rs, feed: extension}

	tmp, err := template.New("css").Funcs(helpers).Funcs(template.FuncMap{
		"extend": rsc.extend,
	}).Parse(rules)

	if err != nil {
		panic(err)
	}

	rsc.template = tmp
	return rsc
}

// Plain returns a new instance of a Rule which uses the raw rule string instead
// of parsing with a template has the source of the stylesheet to be parsed. No processing
// will be done on it.
func Plain(rule string, extension *Rule, rs ...*Rule) *Rule {
	rsc := &Rule{
		plain:   rule,
		feed:    extension,
		depends: rs,
	}

	return rsc
}

// UseExtension sets the css.Rule to be used for extensions and
// returns the rule.
func (r *Rule) UseExtension(c *Rule) *Rule {
	if c == nil {
		return r
	}

	r.feed = c
	return r
}

// Add adds the giving rule into the rules depends list.
func (r *Rule) Add(c *Rule) *Rule {
	r.depends = append(r.depends, c)
	return r
}

// extend attempts to pull a giving set of classes and assigns into
// a target class.
func (r *Rule) extend(item string) string {
	var attrs []string

	for _, rule := range baseStyles.Rules {
		if rule.Prelude != item {
			continue
		}

		for _, prop := range rule.Declarations {
			if prop.Important {
				attrs = append(attrs, prop.StringWithImportant(prop.Important))
			} else {
				attrs = append(attrs, prop.String())
			}
		}
		break
	}

	if len(attrs) == 0 && r.feedStyle != nil {
		for _, rule := range r.feedStyle.Rules {
			if rule.Prelude != item {
				continue
			}

			for _, prop := range rule.Declarations {
				if prop.Important {
					attrs = append(attrs, prop.StringWithImportant(prop.Important))
				} else {
					attrs = append(attrs, prop.String())
				}
			}
			break
		}
	}

	return strings.Join(attrs, "\n")
}

// Stylesheet returns the provided styles using the binding as the argument for the
// provided css template.
func (r *Rule) Stylesheet(bind interface{}, parentNode string) (*bcss.Stylesheet, error) {
	if r.feed != nil {
		sheet, err := r.feed.Stylesheet(bind, parentNode)
		if err != nil {
			return nil, err
		}

		r.feedStyle = sheet
	}

	var stylesheet bcss.Stylesheet

	{
		for _, rule := range r.depends {
			sheet, err := rule.Stylesheet(bind, parentNode)
			if err != nil {
				return nil, err
			}

			stylesheet.Rules = append(stylesheet.Rules, sheet.Rules...)
		}
	}

	var content bytes.Buffer

	if r.template != nil {
		if err := r.template.Execute(&content, bind); err != nil {
			return nil, err
		}
	} else {
		content.WriteString(r.plain)
	}

	sheet, err := parser.Parse(content.String())
	if err != nil {
		return nil, err
	}

	for _, rule := range sheet.Rules {
		r.morphRule(rule, parentNode)
	}

	stylesheet.Rules = append(stylesheet.Rules, sheet.Rules...)

	return &stylesheet, nil
}

// adjustName adjust the provided name according to the set rules of for specific
// css selectors.
func (r *Rule) adjustName(sel string, parentNode string) string {
	sel = strings.TrimSpace(sel)

	switch {
	case strings.Contains(sel, "&"):
		return strings.Replace(sel, "&", parentNode, -1)

	case strings.HasPrefix(sel, ":"):
		return parentNode + "" + sel

	default:
		return sel
	}
}

// morphRules adjusts the provided rules with the parent selector.
func (r *Rule) morphRule(base *bcss.Rule, parentNode string) {
	for index, sel := range base.Selectors {
		base.Selectors[index] = r.adjustName(sel, parentNode)
	}

	for _, rule := range base.Rules {
		if rule.Kind == bcss.AtRule {
			r.morphRule(rule, parentNode)
			continue
		}

		for index, sel := range rule.Selectors {
			rule.Selectors[index] = r.adjustName(sel, parentNode)
		}
	}
}
