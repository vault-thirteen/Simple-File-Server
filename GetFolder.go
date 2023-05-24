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

	absFolderPath := filepath.Join(sfs.rootFolderPath, relPath)

	var absFilePath string
	for _, fdf := range sfs.folderDefaultFiles {
		absFilePath = filepath.Join(absFolderPath, fdf)

		if sfs.isCachingEnabled {
			bytes, fileExists, err = sfs.getDefaultFileUsingCache(absFilePath)
		} else {
			bytes, fileExists, err = sfs.getDefaultFileWithoutCache(absFilePath)
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

// getDefaultFileUsingCache returns contents and a flag showing existence of the
// default file using cache. Path is absolute.
func (sfs *SimpleFileServer) getDefaultFileUsingCache(absFilePath string) (bytes []byte, fileExists bool, err error) {
	fileExists, err = sfs.getFileExistenceUsingCache(absFilePath)
	if err != nil {
		return nil, false, err
	}
	if !fileExists {
		return nil, false, nil
	}

	bytes, err = sfs.cache.GetRecord(absFilePath)
	if err == nil {
		// File is cached. Everything is good.
		return bytes, true, nil
	}

	// File is not cached. Save the file into cache.
	bytes, err = ReadFileFromOs(absFilePath)
	if err != nil {
		return nil, true, err
	}

	err = sfs.cache.AddRecord(absFilePath, bytes)
	if err != nil {
		return nil, true, err
	}

	return bytes, true, nil
}

// getDefaultFileWithoutCache returns contents and a flag showing existence of
// the default file not using cache. Path is absolute.
func (sfs *SimpleFileServer) getDefaultFileWithoutCache(absFilePath string) (bytes []byte, fileExists bool, err error) {
	return sfs.getFileWithoutCache(absFilePath)
}
