package packers_test

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/influx6/trail/assets"
	"github.com/influx6/trail/assets/packers"
	"github.com/influx6/faux/tests"
)

var (
	gopath  = os.Getenv("GOPATH")
	thisSrc = filepath.Join(gopath, "src/github.com/influx6/trail")
)

func TestLessPacker(t *testing.T) {
	expected := "header {\n  color: red;\n  font-size: 24px;\n}\n"
	fixtures := filepath.Join(thisSrc, "assets/packers/fixtures")
	bombless := filepath.Join(fixtures, "bomb.less")
	bomblessRel := filepath.Join("./packers/fixtures/", "bomb.less")

	var less packers.LessPacker

	response, err := less.Pack([]assets.FileStatement{{
		Path:    bomblessRel,
		AbsPath: bombless,
	}}, assets.DirStatement{})

	if err != nil {
		tests.Failed("Should have successfully packed less file: %+q", err)
	}
	tests.Passed("Should have successfully packed less file")

	if len(response) != 1 {
		tests.Failed("Should have successfully received converted less file")
	}
	tests.Passed("Should have successfully received converted less file")

	var b bytes.Buffer
	if _, err := response[0].Writer.WriteTo(&b); err != nil {
		tests.Failed("Should have successfully written data to buffer: %+q", err)
	}
	tests.Passed("Should have successfully written data to buffer")

	if b.String() != expected {
		tests.Failed("Should have successfully matched css output with expected")
	}
	tests.Passed("Should have successfully matched css output with expected")
}
