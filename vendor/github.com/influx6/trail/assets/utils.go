// +build !js

package assets

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

//===============================================================================

// FileStatement defines a structure to contain the details of a given file.
type FileStatement struct {
	Path    string `json:"path"`
	AbsPath string `json:"abspath"`
}

// DirStatement defines a structure for all files retrieved from the giving directory.
type DirStatement struct {
	Total      int
	DirRoot    string
	FilesByExt map[string][]FileStatement
}

//===============================================================================

// GetDirStatement returns a instance of a DirStatement which contains all files
// retrieved through running the directory.
func GetDirStatement(dir string, doGo bool) (DirStatement, error) {
	var statement DirStatement
	statement.FilesByExt = make(map[string][]FileStatement, 0)

	return statement, WalkDir(dir, func(relPath string, absolutePath string, info os.FileInfo) bool {
		if info.IsDir() {
			return true
		}

		if strings.Contains(absolutePath, ".git") {
			return true
		}

		ext := getExtension(relPath)

		if !doGo && ext == ".go" {
			return true
		}

		fileStatement := FileStatement{
			Path:    relPath,
			AbsPath: absolutePath,
		}

		if sets, ok := statement.FilesByExt[ext]; ok {
			statement.FilesByExt[ext] = append(sets, fileStatement)
			return true
		}

		statement.FilesByExt[ext] = []FileStatement{fileStatement}

		return true
	})
}

//===============================================================================

var errStopWalking = errors.New("stop walking directory")

// DirWalker defines a function type which for processing a path and it's info
// retrieved from the fs.
type DirWalker func(rel string, abs string, info os.FileInfo) bool

// WalkDir will run through the provided path which is expected to be a directory
// and runs the provided callback with the current path and FileInfo.
func WalkDir(dir string, callback DirWalker) error {
	isWin := runtime.GOOS == "windows"

	cerr := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		// If we got an error then stop and return it.
		if err != nil {
			return err
		}

		// If its a symlink, don't deal with it.
		if !info.Mode().IsRegular() {
			return nil
		}

		// If on windows, correct path slash.
		if isWin {
			path = filepath.ToSlash(path)
		}

		// Retrive relative path for giving path.
		relPath, err := filepath.Rel(dir, path)
		if err != nil {
			return err
		}

		// If false is return then stop walking and return errStopWalking.
		if !callback(relPath, path, info) {
			return errStopWalking
		}

		return nil
	})

	// If we received error to stop walking then skip
	if cerr == errStopWalking {
		return nil
	}

	return cerr
}

// ParseDirWithExt returns a new instance of all CSS files located within the provided directory.
func ParseDirWithExt(dir string, allowedExtensions []string) (map[string]string, error) {
	items := make(map[string]string)

	// Walk directory pulling contents into css items.
	if cerr := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if cerr := walkDir(allowedExtensions, items, dir, path, info, err); cerr != nil {
			return cerr
		}

		return nil
	}); cerr != nil {
		return nil, cerr
	}

	return items, nil
}

//===============================================================================

// getExtension will return extension associated with a file and ensures to take
// care of files with multiple extension periods, like static.html, .tml.html, as
// well as single extensions suffix like .html, .eml.
func getExtension(name string) string {
	exts := strings.SplitN(name, ".", 2)
	if len(exts) < 2 {
		return ""
	}

	return "." + exts[1]
}

// validExension returns true/false if the extension provide is a valid acceptable one
// based on the allowedExtensions string slice.
func validExtension(extensions []string, ext string) bool {
	for _, es := range extensions {
		if es != ext {
			continue
		}

		return true
	}

	return false
}

// walkDir adds the giving path if it matches certain criterias into the items map.
func walkDir(extensions []string, items map[string]string, root string, path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if !info.Mode().IsRegular() {
		return nil
	}

	// Is file an exension we allow else skip.
	if len(extensions) != 0 && !validExtension(extensions, filepath.Ext(path)) {
		return nil
	}

	rel, err := filepath.Rel(root, path)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	items[rel] = string(data)
	return nil
}
