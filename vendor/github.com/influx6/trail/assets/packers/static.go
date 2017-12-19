// +build !js

package packers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/gu-io/gu/trees"
	"github.com/influx6/faux/fmtwriter"
	"github.com/influx6/moz/gen"
	"github.com/influx6/trail/assets"
	"github.com/influx6/trail/generators/data"
)

// StaticMarkupPacker defines a struct which implements the assets.Packer interface
// and will convert all .static files into go files with the file html content
// turned into type-safe trees.Markup structures(see github.com/influx6/trail/tree/master/trees).
type StaticMarkupPacker struct {
	PackageName     string
	DestinationFile string
}

// Pack process all '.static.html' files present in the FileStatment slice and returns WriteDirectives
// which conta ins expected outputs for these files.
func (static StaticMarkupPacker) Pack(statements []assets.FileStatement, dir assets.DirStatement) ([]assets.WriteDirective, error) {
	blocks := make(map[string]string, 0)

	for _, statement := range statements {
		fileHTML, err := ioutil.ReadFile(statement.AbsPath)
		if err != nil {
			return nil, fmt.Errorf("Failed to read file %q: %s", statement.AbsPath, err)
		}

		writer, err := trees.ParseTreeToText(string(fileHTML), true)
		if err != nil {
			return nil, fmt.Errorf("Failed to parse markup: %q", err.Error())
		}

		var treeWriter bytes.Buffer
		if _, err := writer.WriteTo(&treeWriter); err != nil {
			return nil, fmt.Errorf("Failed to write markup to code for file %q: %s", statement.AbsPath, err)
		}

		blocks[statement.Path] = treeWriter.String()
	}

	return []assets.WriteDirective{
		{
			Static: &assets.StaticDirective{
				WriteInFile: true,
				FileName:    filepath.Base(static.DestinationFile),
				DirName:     filepath.Dir(static.DestinationFile),
			},
			Writer: fmtwriter.New(gen.Block(
				gen.SourceText(
					string(data.Must("trees.gen")),
					struct {
						Trees   map[string]string
						Package string
					}{
						Trees:   blocks,
						Package: static.PackageName,
					},
				),
			), true, true),
			OriginPath:    static.DestinationFile,
			OriginAbsPath: static.DestinationFile,
		},
	}, nil

}
