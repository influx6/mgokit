// +build ignore

package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/influx6/moz/gen"
	"github.com/influx6/trail/assets"
	"github.com/influx6/trail/assets/packers"
)

var (
	version    = "0.0.1" // rely on linker -ldflags -X main.version"
	gitCommit  = ""      // rely on linker: -ldflags -X main.gitCommit"
	getVersion = *flag.Bool("v", false, "Print version")
)

func main() {
	flag.Usage = printUsage
	flag.Parse()

	// if we are to print getVersion.
	if getVersion {
		printVersion()
		return
	}

	publicBundle()

	log.Println("Done!")
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
	fmt.Fprintf(os.Stdout, `Usage: go run [file.gp] [command]

COMMANDS:

	public						# Generate all asset bundle for ./public files
	settings					# Generate asset files from settings

FLAGS:
  -v          Print version.

`)
}

func publicBundle() {
	aspacker := assets.New(packers.RawPacker{})

	aspacker.Register(".js", packers.JSPacker{})
	aspacker.Register(".js.map", packers.JSPacker{})

	aspacker.Register(".css", packers.CSSPacker{CleanCSS: true})
	aspacker.Register(".static.html", packers.StaticMarkupPacker{
		PackageName:     "static",
		DestinationFile: ".//static_bundle.go",
	})

	writer, statics, err := aspacker.Compile("./", false)
	if err != nil {
		log.Fatalf("Failed to get compile asset list: %+q", err)
		return
	}

	pipeGen := gen.Block(
		gen.Package(
			gen.Name("static"),
			writer,
		),
	)

	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current directory: %+q", err)
		return
	}

	if err := writeToFile(pipeGen, "bundle.go", "./", currentDir); err != nil {
		log.Fatalf("Failed to write file: %+q", err)
		return
	}

	for _, directives := range statics {
		for _, directive := range directives {
			if directive.Static == nil {
				continue
			}

			if err := writeToFile(directive.Writer, directive.Static.FileName, directive.Static.DirName, currentDir); err != nil {
				log.Fatalf("Failed to write file: %+q", err)
				return
			}
		}
	}

	log.Println("Bundling completed for 'static'")
}

// writeToFile writes the giving content from the WriterTo instance to the file of
// the giving file.
func writeToFile(w io.WriterTo, fileName string, dirName string, currentDir string) error {
	coDir := filepath.Join(currentDir, dirName)

	if dirName != "" {
		if _, err := os.Stat(coDir); err != nil {
			if err := os.MkdirAll(coDir, 0700); err != nil && err != os.ErrExist {
				return err
			}

			fmt.Printf("- Created package directory: %q\n", coDir)
		}
	}

	coFile := filepath.Join(coDir, fileName)
	file, err := os.Create(coFile)
	if err != nil {
		return err
	}

	defer file.Close()

	if _, err := w.WriteTo(file); err != nil {
		return err
	}

	fmt.Printf("- Created directory file: %q\n", filepath.Join(dirName, fileName))
	return nil
}
