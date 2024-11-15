package sfs

import (
	"strings"
)

// IsPathValid checks validity of the path.
func IsPathValid(path string) bool {
	if strings.Contains(path, PathLevelUp) {
		return false
	}

	return true
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
