package sfs

import (
	"errors"
	"path/filepath"
)

// GetFolder returns contents of the first found default file of a folder and a
// flag showing existence of the default file. Path is a relative path inside
// the file server's root folder.
func (sfs *SimpleFileServer) GetFolder(relPath string) (bytes []byte, fileExists bool, err error) {
	if len(sfs.folderDefaultFiles) == 0 {
		return nil, false, nil
	}

	if !IsPathValid(relPath) {
		return nil, false, errors.New(ErrPathIsNotValid)
	}

	var relFilePath string
	for _, fdf := range sfs.folderDefaultFiles {
		relFilePath = filepath.Join(relPath, fdf)

		if sfs.isCachingEnabled {
			bytes, fileExists, err = sfs.getFileUsingCache(relFilePath)
		} else {
			bytes, fileExists, err = sfs.getFileWithoutCache(relFilePath)
		}
		if err != nil {
			return nil, false, err
		}
		if fileExists {
			return bytes, true, nil
		}
	}

	return nil, false, nil
}
