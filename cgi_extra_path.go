package sfs

import (
	"errors"
	"path/filepath"
	"strings"

	"github.com/vault-thirteen/auxie/file"
)

const (
	ErrCgiExtraPathIsNotFound = "CGI extra path is not found"
)

const (
	ForwardSlashString  = string(`/`)
	BackwardSlashString = string(`\`)
)

// FindExtraPath tries to find CGI extra path.
// "Extra Path" is a CGI interface feature that allows to make crazy-looking
// URLs which are impossible to be parsed, something like the following:
// http://some.machine/cgi-bin/display.pl/cgi/cgi_doc.txt
func (sfs *SimpleFileServer) FindExtraPath(relPath string) (path string, extraPath string, err error) {
	curPath := strings.TrimSuffix(relPath, ForwardSlashString)
	curPath = strings.TrimSuffix(curPath, BackwardSlashString)

	var fileExists bool
	var absFilePath string

	for {
		// Check.
		absFilePath = filepath.Join(sfs.rootFolderPath, curPath)
		fileExists, err = sfs.getFileExistenceWithoutCache(absFilePath)
		if (err != nil) && (err.Error() != file.ErrObjectIsNotFile) {
			return "", "", err
		}
		if fileExists {
			extraPath = strings.TrimPrefix(relPath, curPath)
			return curPath, extraPath, nil
		}

		// Next.
		curPath = shortenPath(curPath)
		if len(curPath) == 0 {
			break
		}
	}

	return "", "", errors.New(ErrCgiExtraPathIsNotFound)
}

// shortenPath trims the last segment of the path.
func shortenPath(path string) (shortenedPath string) {
	dir, _ := filepath.Split(path)
	shortenedPath = strings.TrimSuffix(dir, ForwardSlashString)
	shortenedPath = strings.TrimSuffix(shortenedPath, BackwardSlashString)
	return shortenedPath
}

// GetAbsolutePath returns an absolute path from a relative one.
func (sfs *SimpleFileServer) GetAbsolutePath(relPath string) (absPath string) {
	return filepath.Join(sfs.rootFolderPath, relPath)
}
