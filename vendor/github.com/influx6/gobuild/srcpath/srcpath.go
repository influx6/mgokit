package srcpath

import (
	"os"
	"path/filepath"
)

var (
	goPath        = os.Getenv("GOPATH")
	goRoot        = os.Getenv("GOROOT")
	goSrcPath     = filepath.Join(goPath, "src")
	goRootSrcPath = filepath.Join(goRoot, "src")
)

// RootPath returns current go src path.
func RootPath() string {
	return goRoot
}

// SrcPath returns current go src path.
func SrcPath() string {
	return goSrcPath
}

// FromRootPath returns the giving path as absolute from the GOROOT path
// where the internal packages are stored.
func FromRootPath(pr string) string {
	return filepath.Join(goRootSrcPath, pr)
}

// FromSrcPath returns the giving path as absolute from the gosrc path.
func FromSrcPath(pr string) string {
	return filepath.Join(goSrcPath, pr)
}

// RelativeToRoot returns a path that is relative to GOROOT/src path.
// Where GOROOT, is where the go runtime src is located.
func RelativeToRoot(path string) (string, error) {
	return filepath.Rel(goRootSrcPath, path)
}

// RelativeToSrc returns a path that is relative to the go src path.
func RelativeToSrc(path string) (string, error) {
	return filepath.Rel(goSrcPath, path)
}
