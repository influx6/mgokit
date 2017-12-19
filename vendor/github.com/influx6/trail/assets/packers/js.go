// +build !js

package packers

import (
	"bytes"
	"io"
	"os"

	"github.com/influx6/trail/assets"
)

// JSPacker defines an implementation for parsing css files.
type JSPacker struct {
	Exceptions []string
}

// Pack process all files present in the FileStatment slice and returns WriteDirectives
// which contains expected outputs for these files.
func (less JSPacker) Pack(statements []assets.FileStatement, dir assets.DirStatement) ([]assets.WriteDirective, error) {
	var directives []assets.WriteDirective

	for _, statement := range statements {
		// Validate that we do not have the relative or absolute path as exceptions.
		if less.hasException(statement.Path) || less.hasException(statement.AbsPath) {
			continue
		}

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

// hasException validates the path is not within the exception list.
func (less JSPacker) hasException(path string) bool {
	for _, pl := range less.Exceptions {
		if pl == path {
			return true
		}
	}

	return false
}
