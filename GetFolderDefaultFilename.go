package sfs

import (
	"errors"
	"path/filepath"
)

// GetFolderDefaultFilename returns name of the first found default file of a
// folder and a flag showing existence of the default file. Path is a relative
// path inside the file server's root folder.
func (sfs *SimpleFileServer) GetFolderDefaultFilename(folderRelPath string) (fileName string, err error) {
	if len(sfs.folderDefaultFiles) == 0 {
		return "", nil
	}

	if !IsPathValid(folderRelPath) {
		return "", errors.New(ErrPathIsNotValid)
	}

	var fileRelPath string
	var fileExists bool
	for _, fdf := range sfs.folderDefaultFiles {
		fileRelPath = filepath.Join(folderRelPath, fdf)

		fileExists, err = sfs.FileExists(fileRelPath)
		if err != nil {
			return "", err
		}
		if fileExists {
			return fdf, nil
		}
	}

	return "", errors.New(ErrFileIsNotFound)
}
