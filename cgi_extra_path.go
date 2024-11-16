package sfs

import (
	"errors"
	"path/filepath"
	"strings"

	"github.com/vault-thirteen/auxie/file"
)

// FindExtraPath tries to find CGI extra path.
// "Extra Path" is a CGI interface feature that allows to make crazy-looking
// URLs which are impossible to be parsed, something like the following:
// http://some.machine/cgi-bin/display.pl/cgi/cgi_doc.txt
func (sfs *SimpleFileServer) FindExtraPath(relPath string) (path string, extraPath string, err error) {
	curPath := strings.TrimSuffix(relPath, ForwardSlashString)
	curPath = strings.TrimSuffix(curPath, BackwardSlashString)

	var fileExists bool

	for {
		// Check.
		fileExists, err = sfs.fileExistsInStorage(curPath)
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

	return "", "", errors.New(Err_CgiExtraPathIsNotFound)
}

// shortenPath trims the last segment of the path.
func shortenPath(path string) (shortenedPath string) {
	dir, _ := filepath.Split(path)
	shortenedPath = strings.TrimSuffix(dir, ForwardSlashString)
	shortenedPath = strings.TrimSuffix(shortenedPath, BackwardSlashString)
	return shortenedPath
}
