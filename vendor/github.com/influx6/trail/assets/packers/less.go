// +build !js

package packers

import (
	"bytes"
	"context"
	"fmt"
	"os"
	gexec "os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/influx6/trail/assets"
	"github.com/influx6/faux/exec"
	"github.com/influx6/faux/metrics"
)

var (
	lessBin = filepath.Join(inGOPATHSrc, "github.com/influx6/trail/node_modules/less/bin")
)

// LessPacker defines an implementation for parsing .less files into css files using the less compiler in nodejs.
// WARNING: Requires Nodejs to be installed.
type LessPacker struct {
	MainFile string
	Options  map[string]string
}

// Pack process all files present in the FileStatment slice and returns WriteDirectives
// which conta ins expected outputs for these files.
func (less LessPacker) Pack(statements []assets.FileStatement, dir assets.DirStatement) ([]assets.WriteDirective, error) {
	var directives []assets.WriteDirective

	// If main less file has being set then attempt to find main file.
	if less.MainFile == "" {
		for _, statement := range statements {
			if err := processStatement(statement, less, &directives); err != nil {
				return nil, err
			}
		}

		return directives, nil
	}

	for _, statement := range statements {
		if statement.Path != less.MainFile {
			continue
		}

		if err := processStatement(statement, less, &directives); err != nil {
			return nil, err
		}
	}

	return directives, nil
}

func processStatement(statement assets.FileStatement, less LessPacker, directives *[]assets.WriteDirective) error {
	fileExt := filepath.Ext(statement.Path)
	cssFileName := filepath.Join(filepath.Dir(statement.Path), strings.Replace(filepath.Base(statement.Path), fileExt, ".css", 1))
	cssAbsFileName := filepath.Join(filepath.Dir(statement.AbsPath), strings.Replace(filepath.Base(statement.Path), fileExt, ".css", 1))

	cssFileName = strings.Replace(cssFileName, "less/", "css/", 1)
	cssAbsFileName = strings.Replace(cssAbsFileName, "less/", "css/", 1)

	var args []string

	for option, value := range less.Options {
		args = append(args, option, value)
	}

	args = append(args, filepath.Clean(statement.AbsPath))

	node, err := gexec.LookPath("node")
	if err != nil {
		return err
	}

	os.Setenv("node", node)

	command := fmt.Sprintf("%s %s", filepath.Join(lessBin, "lessc"), strings.Join(args, " "))

	var errBuf, outBuf bytes.Buffer
	cleanCmd := exec.New(
		exec.Async(),
		exec.Command(command),
		exec.Output(&outBuf),
		exec.Err(&errBuf),
	)

	ctx, cancl := context.WithTimeout(context.Background(), time.Minute)
	defer cancl()

	if err := cleanCmd.Exec(ctx, metrics.New()); err != nil {
		return fmt.Errorf("Command Execution Failed: %+q\n Response: %+q\n Command: %+q", err, errBuf.String(), command)
	}

	*directives = append(*directives, assets.WriteDirective{
		OriginPath:    cssFileName,
		OriginAbsPath: cssAbsFileName,
		Writer:        bytes.NewReader(outBuf.Bytes()),
	})

	return nil
}
