package jrutil

import (
	"path/filepath"
)

// Realpath returns the real path given a file name or an error.
func Realpath(p string) (string, error) {
	var err  error
	var aPath string
	var rPath string

	// Make sure we have an absolute path.
	aPath, err = filepath.Abs(p)
	if err != nil {
		return "", err
	}

	// Dereference the symlinks.
	rPath, err = filepath.EvalSymlinks(aPath)
	if err != nil {
		return "", err
	}

	return rPath, nil
}
