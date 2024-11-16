package sfs

import (
	"errors"

	"github.com/vault-thirteen/auxie/file"
)

func (sfs *SimpleFileServer) GetFile(relPath string) (data []byte, err error) {
	if !IsPathValid(relPath) {
		return nil, errors.New(Err_PathIsNotValid)
	}

	if !sfs.isCachingEnabled {
		return sfs.getFileFromStorage(relPath)
	}

	if sfs.fileExistsInCache(relPath) {
		return sfs.getFileFromCache(relPath)
	}

	data, err = sfs.getFileFromStorage(relPath)
	if err != nil {
		return nil, err
	}

	err = sfs.createFileInCache(relPath, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (sfs *SimpleFileServer) getFileFromCache(relPath string) (bytes []byte, err error) {
	return sfs.cache.GetRecord(relPath)
}

func (sfs *SimpleFileServer) getFileFromStorage(relPath string) (bytes []byte, err error) {
	return file.GetFileContents(sfs.GetAbsolutePath(relPath))
}
