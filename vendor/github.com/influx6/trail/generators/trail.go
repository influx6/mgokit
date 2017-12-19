// +build !js

package generators

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/influx6/faux/fmtwriter"

	"github.com/influx6/moz/ast"
	"github.com/influx6/moz/gen"
	"github.com/influx6/trail/generators/data"
)

var (
	inGOPATH    = os.Getenv("GOPATH")
	inGOPATHSrc = filepath.Join(inGOPATH, "src")
	badSymbols  = regexp.MustCompile(`[(|\-|_|\W|\d)+]`)
	notAllowed  = regexp.MustCompile(`[^(_|\w|\d)+]`)
)

// TrailPackages returns a slice of WriteDirectives which contain data to be written to disk to create
// a suitable package for asset bundle.
func TrailPackages(an ast.AnnotationDeclaration, pkg ast.PackageDeclaration, pkgDir ast.Package) ([]gen.WriteDirective, error) {
	if len(an.Arguments) == 0 {
		return nil, errors.New("Expected atleast one argument for annotation as component name")
	}

	workDir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve current directory path: %+q", err)
	}

	gridCSSData := data.Must("grid.css.gen")
	flashCSSData := data.Must("flash.css.gen")
	baseCSSData := data.Must("base.css.gen")
	gridNormCSS := data.Must("normalize.css.gen")

	packageDir, err := filepath.Rel(inGOPATHSrc, workDir)
	if err != nil {
		fmt.Printf("Failed to retrieve package directory path in go src: %+q\n", err)
	}

	componentName := badSymbols.ReplaceAllString(an.Arguments[0], "")

	var targetDir string

	if componentName != "" {
		targetDir = strings.ToLower(componentName)
	} else {
		componentName = filepath.Base(workDir)
	}

	componentNameLower := strings.ToLower(componentName)
	componentPackageDir := filepath.Join(packageDir, targetDir)

	publicStandInGen := gen.Block(
		gen.Package(
			gen.Name(componentNameLower),
			gen.SourceText(
				string(data.Must("bundle-standin.gen")),
				struct {
					Name    string
					Package string
				}{
					Name:    componentName,
					Package: componentNameLower,
				},
			),
		),
	)

	publicGen := gen.Block(
		gen.SourceText(
			string(data.Must("pack-bundle.gen")),
			struct {
				Name          string
				LessFile      string
				Package       string
				TargetDir     string
				TargetPackage string
				Settings      bool
			}{
				TargetDir:     "./",
				TargetPackage: componentNameLower,
				Settings:      true,
				Name:          componentName,
				Package:       componentNameLower,
				LessFile:      fmt.Sprintf("less/%s.less", componentNameLower),
			},
		),
	)

	lessGen := gen.Block(
		gen.SourceText(
			string(data.Must("main.less.gen")),
			struct{}{},
		),
	)

	jsGen := gen.Block(gen.SourceText(string(data.Must("jquery.min.js.gen")), struct{}{}))

	tomlGen := gen.Block(
		gen.SourceText(
			string(data.Must("settings.toml.gen")),
			struct {
				Name    string
				Package string
			}{
				Name:    componentNameLower,
				Package: componentPackageDir,
			},
		),
	)

	docGen := gen.Block(
		gen.Package(
			gen.Name(componentNameLower),
			gen.Block(
				gen.Text("\n"),
				gen.Text("//go:generate go run generate.go settings"),
				gen.Text("\n"),
				gen.Text("//go:generate go run generate.go public"),
				gen.Text("\n"),
			),
		),
	)

	lessName := "index"
	if componentName != "" {
		lessName = componentNameLower
	}

	commands := []gen.WriteDirective{
		{
			DontOverride: false,
			Dir:          targetDir,
		},
		{
			DontOverride: true,
			FileName:     "index.html",
			Dir:          filepath.Join(targetDir, "layout"),
			Writer:       bytes.NewBuffer(bytes.Replace(data.Must("base.html.gen"), []byte("{{.Name}}"), []byte(componentName), -1)),
		},
		{
			DontOverride: true,
			Writer:       docGen,
			FileName:     "doc.go",
			Dir:          targetDir,
		},
		{
			DontOverride: true,
			FileName:     "main.js",
			Dir:          filepath.Join(targetDir, "js"),
			Writer:       bytes.NewBufferString("//strictmode"),
		},
		{
			DontOverride: false,
			Writer:       jsGen,
			FileName:     "jquery.min.js",
			Dir:          filepath.Join(targetDir, "js"),
		},
		{
			DontOverride: true,
			Dir:          filepath.Join(targetDir, "js"),
			FileName:     "flash.js",
			Writer:       bytes.NewBuffer(data.Must("flash.js.gen")),
		},
		{
			DontOverride: false,
			Dir:          filepath.Join(targetDir, "css"),
			FileName:     "normalize.css",
			Writer:       bytes.NewBuffer(gridNormCSS),
		},
		{
			DontOverride: true,
			Dir:          filepath.Join(targetDir, "css"),
			FileName:     "flash.css",
			Writer:       bytes.NewBuffer(flashCSSData),
		},
		{
			DontOverride: false,
			Dir:          filepath.Join(targetDir, "css"),
			FileName:     "grid.css",
			Writer:       bytes.NewBuffer(gridCSSData),
		},
		{
			DontOverride: true,
			Dir:          filepath.Join(targetDir, "css"),
			FileName:     "base.css",
			Writer:       bytes.NewBuffer(baseCSSData),
		},
		{
			DontOverride: false,
			Dir:          filepath.Join(targetDir, "less"),
		},
		{
			DontOverride: true,
			Writer:       lessGen,
			Dir:          filepath.Join(targetDir, "less"),
			FileName:     fmt.Sprintf("%s.less", lessName),
		},
		{
			DontOverride: true,
			Writer:       tomlGen,
			Dir:          targetDir,
			FileName:     "settings.toml",
		},
		{
			DontOverride: false,
			Dir:          targetDir,
			FileName:     "generate.go",
			Writer:       fmtwriter.New(publicGen, true, true),
		},
		{
			DontOverride: false,
			Dir:          targetDir,
			FileName:     "bundle.go",
			Writer:       fmtwriter.New(publicStandInGen, true, true),
		},
	}

	return commands, nil
}

func validateName(val string) bool {
	return notAllowed.MatchString(val)
}
