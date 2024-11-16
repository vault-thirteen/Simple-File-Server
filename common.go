package sfs

import (
	"path/filepath"
	"strings"
)

const (
	ForwardSlashString  = string(`/`)
	BackwardSlashString = string(`\`)
	PathLevelUp         = ".."
)

const (
	Err_CgiExtraPathIsNotFound = "CGI extra path is not found"
	Err_ActionIsImpossible     = "action is impossible"
	Err_CacheIsDisabled        = "cache is disabled"
	Err_FolderDoesNotExist     = "folder does not exist"
	Err_PathIsNotValid         = "path is not valid"
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

// GetAbsolutePath returns an absolute path from a relative one.
func (sfs *SimpleFileServer) GetAbsolutePath(relPath string) (absPath string) {
	return filepath.Join(sfs.rootFolderPath, relPath)
}
