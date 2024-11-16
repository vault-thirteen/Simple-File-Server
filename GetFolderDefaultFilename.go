package sfs

import (
	"errors"
	"path/filepath"
)

// GetFolderDefaultFilename returns a first found default file name for the
// folder. If no default files are set, an empty string is returned. If default
// files are set, but none is found, an empty string is returned.
func (sfs *SimpleFileServer) GetFolderDefaultFilename(folderRelPath string) (fileName string, err error) {
	if len(sfs.folderDefaultFiles) == 0 {
		return "", nil
	}

	if !IsPathValid(folderRelPath) {
		return "", errors.New(Err_PathIsNotValid)
	}

	if !sfs.isCachingEnabled {
		return sfs.getFolderDefaultFilenameFromStorage(folderRelPath)
	}

	return sfs.getFolderDefaultFilenameFromStorage(folderRelPath)
}

func (sfs *SimpleFileServer) getFolderDefaultFilenameFromCache(folderRelPath string) (fileName string, err error) {
	return "", errors.New(Err_ActionIsImpossible)
}

func (sfs *SimpleFileServer) getFolderDefaultFilenameFromStorage(folderRelPath string) (fileName string, err error) {
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

	return "", nil
}
