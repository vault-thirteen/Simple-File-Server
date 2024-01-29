package sfs

import (
	"errors"
	"path/filepath"
)

// GetFile returns file's contents and a flag showing existence of the
// requested file. Path is a relative path inside the file server's root
// folder.
func (sfs *SimpleFileServer) GetFile(relPath string) (bytes []byte, fileExists bool, err error) {
	if !IsPathValid(relPath) {
		return nil, false, errors.New(ErrPathIsNotValid)
	}

	if sfs.isCachingEnabled {
		return sfs.getFileUsingCache(relPath)
	} else {
		return sfs.getFileWithoutCache(relPath)
	}
}

// getFileUsingCache returns contents and a flag showing existence of the file
// using cache. Path is absolute.
func (sfs *SimpleFileServer) getFileUsingCache(relPath string) (bytes []byte, fileExists bool, err error) {
	bytes, err = sfs.cache.GetRecord(relPath)
	if err == nil {
		// File is cached. Everything is good.
		return bytes, true, nil
	}

	// File is not cached.
	bytes, fileExists, err = sfs.getFileWithoutCache(relPath)
	if err != nil {
		return nil, false, err
	}
	if !fileExists {
		return nil, false, nil
	}

	err = sfs.cache.AddRecord(relPath, bytes)
	if err != nil {
		return nil, true, err
	}

	return bytes, true, nil
}

// getFileWithoutCache returns contents and a flag showing existence of the file
// not using cache. Path is absolute.
func (sfs *SimpleFileServer) getFileWithoutCache(relPath string) (bytes []byte, fileExists bool, err error) {
	absFilePath := filepath.Join(sfs.rootFolderPath, relPath)

	fileExists, err = sfs.getFileExistenceWithoutCache(absFilePath)
	if err != nil {
		return nil, false, err
	}
	if !fileExists {
		return nil, false, nil
	}

	bytes, err = ReadFileFromOs(absFilePath)
	if err != nil {
		return nil, true, err
	}

	return bytes, true, nil
}
