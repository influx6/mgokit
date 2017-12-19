package generators

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/influx6/faux/fmtwriter"

	"github.com/influx6/moz/ast"
	"github.com/influx6/moz/gen"
	"github.com/influx6/trail/generators/data"
)

// TrailView returns a series of file commands which create asset bundling for a giving file.
func TrailView(an ast.AnnotationDeclaration, pkg ast.PackageDeclaration, pk ast.Package) ([]gen.WriteDirective, error) {
	if len(an.Arguments) == 0 {
		return nil, errors.New("Expected atleast one argument for annotation as component name")
	}

	workDir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve current directory path: %+q", err)
	}

	componentName := badSymbols.ReplaceAllString(an.Arguments[0], "")

	var targetDir string

	if componentName != "" {
		targetDir = strings.ToLower(componentName)
	} else {
		componentName = filepath.Base(workDir)
	}

	componentNameLower := strings.ToLower(componentName)

	generatorGen := gen.Block(
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
				Name:          componentName,
				Package:       componentNameLower,
			},
		),
	)

	pipeGen := gen.Block(
		gen.Package(
			gen.Name(componentNameLower),
			gen.Block(
				gen.Text("\n"),
				gen.SourceText(
					string(data.Must("bundle-standin.gen")),
					nil,
				),
			),
		),
	)

	baseGen := gen.Block(
		gen.Package(
			gen.Name(componentNameLower),
			gen.Block(
				gen.Text("\n"),
				gen.Text("//go:generate go run generate.go"),
				gen.Text("\n"),
			),
		),
	)

	flashTemplate := `<div class="flash block one-whole">
      {{range $name, $messages := flashMessages}}
      {{range $messages}}<div class="one-whole center-aligned hide flash-message {{$name}}">{{.}}</div>{{end}}
      {{end}}</div>`

	htmlGen := gen.Block(
		gen.SourceText(
			string(data.Must("base.html.gen")),
			struct {
				Name          string
				Path          string
				JSFile        string
				FlashTemplate string
			}{
				Name:          componentNameLower,
				Path:          "public",
				JSFile:        fmt.Sprintf("%s/%s", "js", "main.js"),
				FlashTemplate: flashTemplate,
			},
		),
	)

	return []gen.WriteDirective{
		{
			DontOverride: false,
			FileName:     "bundle.go",
			Dir:          targetDir,
			Writer:       fmtwriter.New(pipeGen, true, true),
		},
		{
			DontOverride: true,
			Writer:       htmlGen,
			Dir:          targetDir,
			FileName:     fmt.Sprintf("%s.html", componentNameLower),
		},
		{
			DontOverride: false,
			FileName:     "generate.go",
			Dir:          targetDir,
			Writer:       fmtwriter.New(generatorGen, true, true),
		},
		{
			DontOverride: true,
			FileName:     fmt.Sprintf("%s.go", componentNameLower),
			Dir:          targetDir,
			Writer:       fmtwriter.New(baseGen, true, true),
		},
	}, nil
}

// TrailFiles returns a series of file commands which create asset bundling for a giving file.
func TrailFiles(an ast.AnnotationDeclaration, pkg ast.PackageDeclaration, pk ast.Package) ([]gen.WriteDirective, error) {
	if len(an.Arguments) == 0 {
		return nil, errors.New("Expected atleast one argument for annotation as component name")
	}

	workDir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve current directory path: %+q", err)
	}

	componentName := badSymbols.ReplaceAllString(an.Arguments[0], "")

	var targetDir string

	if componentName != "" {
		targetDir = strings.ToLower(componentName)
	} else {
		componentName = filepath.Base(workDir)
	}

	componentNameLower := strings.ToLower(componentName)

	generatorGen := gen.Block(
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
				Name:          componentName,
				Package:       componentNameLower,
			},
		),
	)

	pipeGen := gen.Block(
		gen.Package(
			gen.Name(componentNameLower),
			gen.Block(
				gen.Text("\n"),
				gen.SourceText(
					string(data.Must("bundle-standin.gen")),
					nil,
				),
			),
		),
	)

	baseGen := gen.Block(
		gen.Package(
			gen.Name(componentNameLower),
			gen.Block(
				gen.Text("\n"),
				gen.Text("//go:generate go run generate.go"),
				gen.Text("\n"),
			),
		),
	)

	return []gen.WriteDirective{
		{
			DontOverride: false,
			FileName:     "bundle.go",
			Dir:          targetDir,
			Writer:       fmtwriter.New(pipeGen, true, true),
		},
		{
			DontOverride: false,
			FileName:     "generate.go",
			Dir:          targetDir,
			Writer:       fmtwriter.New(generatorGen, true, true),
		},
		{
			DontOverride: true,
			FileName:     fmt.Sprintf("%s.go", componentNameLower),
			Dir:          targetDir,
			Writer:       fmtwriter.New(baseGen, true, true),
		},
	}, nil
}
