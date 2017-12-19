// +build ignore

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

var (
	pkg = `// Package style is generated to contain the content of a css template for generating a styleguide for use in projects.

// Document is auto-generate and should not be modified by hand.

//go:generate go run generate.go

package styleguide

// styleTemplate contains the text template used to generated the full set of 
// css template for a giving styleguide.
`

	pkgVar = "var styleTemplate = `%s`\n"
)

func main() {
	js, err := ioutil.ReadFile("./style.sess")
	if err != nil {
		panic(fmt.Sprintf("Unable to locate `style.js` file: %q", err.Error()))
	}

	goFile, err := os.Create("./style_template.go")
	if err != nil {
		panic(fmt.Sprintf("Unable to create `style_template.go` file: %q", err.Error()))
	}

	defer goFile.Close()

	if _, err := fmt.Fprint(goFile, pkg); err != nil {
		panic(fmt.Sprintf("Unable to write data to `style_template.go`: %q", err.Error()))
	}

	if _, err := fmt.Fprintf(goFile, pkgVar, js); err != nil {
		panic(fmt.Sprintf("Unable to write data to `style_template.go`: %q", err.Error()))
	}
}
