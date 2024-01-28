package sfs

import (
	"errors"
	"path/filepath"
)

// GetFolderDefaultFilename returns name of the first found default file of a
// folder and a flag showing existence of the default file. Path is a relative
// path inside the file server's root folder.
func (sfs *SimpleFileServer) GetFolderDefaultFilename(relPath string) (fileName string, fileExists bool, err error) {
	if len(sfs.folderDefaultFiles) == 0 {
		return "", false, nil
	}

	if !IsPathValid(relPath) {
		return "", false, errors.New(ErrPathIsNotValid)
	}

	absFolderPath := filepath.Join(sfs.rootFolderPath, relPath)

	var absFilePath string
	for _, fdf := range sfs.folderDefaultFiles {
		absFilePath = filepath.Join(absFolderPath, fdf)

		fileExists, err = sfs.getFileExistenceWithoutCache(absFilePath)
		if err != nil {
			return "", false, err
		}
		if fileExists {
			return fdf, true, nil
		}
	}

	return "", false, nil
}
