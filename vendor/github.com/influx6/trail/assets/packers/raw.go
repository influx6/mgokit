// +build !js

package packers

import (
	"bytes"
	"io"
	"os"

	"github.com/influx6/trail/assets"
)

// RawPacker defines an implementation for parsing css files.
type RawPacker struct{}

// Pack process all files present in the FileStatment slice and returns WriteDirectives
// which contains expected outputs for these files.
func (less RawPacker) Pack(statements []assets.FileStatement, dir assets.DirStatement) ([]assets.WriteDirective, error) {
	var directives []assets.WriteDirective

	for _, statement := range statements {
		reader, err := os.Open(statement.AbsPath)
		if err != nil {
			return nil, err
		}

		var bu bytes.Buffer
		if _, err := io.Copy(&bu, reader); err != nil && err != io.EOF {
			return nil, err
		}

		directives = append(directives, assets.WriteDirective{
			Writer:        &bu,
			OriginPath:    statement.Path,
			OriginAbsPath: statement.AbsPath,
		})
	}

	return directives, nil
}
