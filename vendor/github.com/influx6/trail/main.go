package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/influx6/moz/ast"
	"github.com/influx6/trail/generators"
)

var (
	version   = "0.0.1" // rely on linker -ldflags -X main.version"
	gitCommit = ""      // rely on linker: -ldflags -X main.gitCommit"
)

var (
	getVersion   = flag.Bool("v", false, "Print version")
	forceRebuild = flag.Bool("f", false, "force rebuild")
)

func main() {
	flag.Usage = printUsage
	flag.Parse()

	// if we are to print getVersion.
	if *getVersion {
		printVersion()
		return
	}

	command := flag.Arg(0)
	name := flag.Arg(1)

	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get directory path: %+q", err)
		return
	}

	switch command {
	case "public":
		generatePublic(currentDir, name)
	case "files":
		generateFiles(currentDir, name)
	case "view":
		generateView(currentDir, name)
	default:
		printUsage()
	}

	log.Println("Trail asset bundling ready!")
}

func generateFiles(currentDir, name string) {
	commands, err := generators.TrailFiles(
		ast.AnnotationDeclaration{Arguments: []string{name}},
		ast.PackageDeclaration{FilePath: currentDir},
		ast.Package{},
	)
	if err != nil {
		log.Fatalf("Failed to generate trail directives: %+q", err)
		return
	}

	if err := ast.SimpleWriteDirectives("", *forceRebuild, commands...); err != nil {
		log.Fatalf("Failed to create package directories: %+q", err)
		return
	}
}

func generateView(currentDir, name string) {
	commands, err := generators.TrailView(
		ast.AnnotationDeclaration{Arguments: []string{name}},
		ast.PackageDeclaration{FilePath: currentDir},
		ast.Package{},
	)
	if err != nil {
		log.Fatalf("Failed to generate trail directives: %+q", err)
		return
	}

	if err := ast.SimpleWriteDirectives("", *forceRebuild, commands...); err != nil {
		log.Fatalf("Failed to create package directories: %+q", err)
		return
	}
}

func generatePublic(currentDir, name string) {
	commands, err := generators.TrailPackages(
		ast.AnnotationDeclaration{Arguments: []string{name}},
		ast.PackageDeclaration{FilePath: currentDir},
		ast.Package{},
	)
	if err != nil {
		log.Fatalf("Failed to generate trail directives: %+q", err)
		return
	}

	if err := ast.SimpleWriteDirectives("", *forceRebuild, commands...); err != nil {
		log.Fatalf("Failed to create package directories: %+q", err)
		return
	}
}

// printVersion prints corresponding build getVersion with associated build stamp and git commit if provided.
func printVersion() {
	fragments := []string{version}

	if gitCommit != "" {
		fragments = append(fragments, fmt.Sprintf("git#%s", gitCommit))
	}

	fmt.Fprint(os.Stdout, strings.Join(fragments, " "))
}

// printUsage prints out usage message for CLI tool.
func printUsage() {
	fmt.Fprintf(os.Stdout, `Usage: trail [options]
Trail creates a package for package of web assets using its internal bundlers.

COMMANDS:

	trail files [optional-name]	# Creates a generate.go file which bundles all assets in created directory.
	trail view [optional-name]	# Creates a generate.go file which bundles all assets in created directory.
	trail public [optional-name]	# Creates a complete package and content for asset bundling all static files

where:

	[optional-name] defines the name for the directory to be used for the assets if provided, else
	having files created within working directory.

EXAMPLES:

	trail view home			# Creates a generate.go file which bundles all assets in create directory.
	trail public static-data	# Creates a complete package and content for asset bundling all static files


FLAGS:
	-v      Print version.
	-f 	Force re-generation of all files
`)
}
