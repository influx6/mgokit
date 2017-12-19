package packers_test

import (
	"bytes"
	"path/filepath"
	"testing"

	"github.com/influx6/trail/assets"
	"github.com/influx6/trail/assets/packers"
	"github.com/influx6/faux/tests"
)

func TestCleanCSSPacker(t *testing.T) {
	expected := "html{width:100%;height:100%}body{width:100%;height:100%}div.tuglife{width:100%;height:100%}"
	fixtures := filepath.Join(thisSrc, "assets/packers/fixtures")
	bombless := filepath.Join(fixtures, "wordan.css")
	bomblessRel := filepath.Join("./packers/fixtures/", "bomb.less")

	clean := packers.CleanCSSPacker{Args: []string{"-O", "1"}}

	response, err := clean.Pack([]assets.FileStatement{{
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
		tests.Info("Expected: %+q", expected)
		tests.Info("Received: %+q", b.String())
		tests.Failed("Should have successfully matched css output with expected")
	}
	tests.Passed("Should have successfully matched css output with expected")
}
