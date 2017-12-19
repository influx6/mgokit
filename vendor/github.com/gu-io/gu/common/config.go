package common

import "errors"

// Remover defines an interface which exposes a remove method.
type Remover interface {
	Remove()
	Add(func())
}

// Settings defines a structure which contains fields that are used to contain
// specific user settings for the gu build system.
type Settings struct {
	App     string `toml:"app"`
	Theme   Theme  `toml:"theme"`
	Package string `toml:"package"`
	Static  Static `toml:"static"`
	Public  Public `toml:"public"`
}

// Validate will validate the state of the giving fields.
func (s Settings) Validate() error {
	if err := s.Public.Validate(); err != nil {
		return err
	}

	return nil
}

// Static defines a configuration which details the place
// where static content and the html file for a driver should be
// located. This will mostly be used by js driver.
type Static struct {
	IndexDir      string `toml:"indexDir"`
	JSFileName    string `toml:"jsFile"`
	JSMapFileName string `toml:"jsMapFile"`
}

// Public defines giving settings for the public assets folder which will be build
// to generate a embeddeable and servable assets package.
type Public struct {
	Path        string `toml:"path"`
	PackageName string `toml:"packageName"`
}

// Validate will validate the state of the giving fields.
func (p Public) Validate() error {
	if p.Path == "" {
		return errors.New("Public.Path must be set")
	}

	if p.PackageName == "" {
		return errors.New("Public.PackageName must be set")
	}

	return nil
}

// Theme defines a struct whhich contains settings for generating a stylesheet of
// css rules.
type Theme struct {
	MinimumScaleCount             int     // Total scale to generate small font sizes.
	MaximumScaleCount             int     // Total scale to generate large font sizes
	MinimumHeadScaleCount         int     // Total scale to generate small font sizes.
	MaximumHeadScaleCount         int     // Total scale to generate large font sizes
	BaseFontSize                  int     // BaseFontSize for typeface using the provide BaseScale.
	SmallBorderRadius             int     // SmallBorderRadius for tiny components eg checkbox, radio buttons.
	MediumBorderRadius            int     // MediaBorderRadius for buttons, inputs, etc
	LargeBorderRadius             int     // LargeBorderRadius for components like cards, modals, etc.
	BaseScale                     float64 // BaseScale to use for generating expansion/detraction scale for font sizes.
	HeaderBaseScale               float64 // BaseScale to use for generating expansion/detraction scale for header h1-h6 tags.
	PrimaryWhite                  string
	SuccessColor                  string
	FailureColor                  string
	PrimaryColor                  string
	HoverShadow                   string
	DropShadow                    string
	BaseShadow                    string
	SecondaryColor                string
	FloatingShadow                string
	PrimaryBrandColor             string
	SecondaryBrandColor           string
	AnimationCurveDefault         string
	AnimationCurveFastOutLinearIn string
	AnimationCurveFastOutSlowIn   string
	AnimationCurveLinearOutSlowIn string
	MaterialPalettes              map[string][]string
}
