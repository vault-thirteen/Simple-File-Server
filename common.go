package sfs

import (
	"io"
	"os"
	"strings"

	"github.com/vault-thirteen/errorz"
)

// IsPathValid checks validity of the path.
func IsPathValid(path string) bool {
	if strings.Contains(path, PathLevelUp) {
		return false
	}
	return true
}

// ReadFileFromOs reads file's contents from an operating system.
func ReadFileFromOs(absFilePath string) (bytes []byte, err error) {
	var f *os.File
	f, err = os.Open(absFilePath)
	if err != nil {
		return nil, err
	}
	defer func() {
		derr := f.Close()
		if derr != nil {
			err = errorz.Combine(err, derr)
		}
	}()

	bytes, err = io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

// IsPathFolder checks whether the specified path is a file or a directory.
func IsPathFolder(path string) bool {
	symbols := []rune(path)
	if len(symbols) == 0 {
		return false
	}

	lastSymbol := symbols[len(symbols)-1]

	if (lastSymbol == '/') || (lastSymbol == '\\') {
		return true
	}

	return false
}
