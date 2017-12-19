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
	inGOPATH         = os.Getenv("GOPATH")
	inGOPATHSrc      = filepath.Join(inGOPATH, "src")
	guSrc            = filepath.Join(inGOPATHSrc, "github.com/influx6/trail")
	guSrcNodeModules = filepath.Join(inGOPATHSrc, "github.com/influx6/trail/node_modules")
	cleanCSSBin      = filepath.Join(inGOPATHSrc, "github.com/influx6/trail/node_modules/clean-css-cli/bin")
)

// CleanCSSPacker defines an implementation for parsing css files.
// WARNING: Requires Nodejs to be installed.
type CleanCSSPacker struct {
	Args []string
}

// Pack process all files present in the FileStatment slice and returns WriteDirectives
// which conta ins expected outputs for these files.
func (cess CleanCSSPacker) Pack(statements []assets.FileStatement, dir assets.DirStatement) ([]assets.WriteDirective, error) {
	var directives []assets.WriteDirective

	for _, statement := range statements {
		if err := processCleanStatement(statement, cess, &directives); err != nil {
			return nil, err
		}
	}

	return directives, nil
}

func processCleanStatement(statement assets.FileStatement, cess CleanCSSPacker, directives *[]assets.WriteDirective) error {
	args := append([]string{}, cess.Args...)
	args = append(args, filepath.Clean(statement.AbsPath))

	node, err := gexec.LookPath("node")
	if err != nil {
		return err
	}

	os.Setenv("node", node)

	command := fmt.Sprintf("%s %s", filepath.Join(cleanCSSBin, "cleancss"), strings.Join(args, " "))

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
		OriginPath:    statement.Path,
		OriginAbsPath: statement.AbsPath,
		Writer:        bytes.NewReader(outBuf.Bytes()),
	})

	return nil
}
