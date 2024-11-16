package sfs

import (
	"errors"

	"github.com/vault-thirteen/auxie/file"
)

func (sfs *SimpleFileServer) FileExists(relPath string) (fileExists bool, err error) {
	if !IsPathValid(relPath) {
		return false, errors.New(Err_PathIsNotValid)
	}

	if !sfs.isCachingEnabled {
		return sfs.fileExistsInStorage(relPath)
	}

	if sfs.fileExistsInCache(relPath) {
		return true, nil
	}

	return sfs.fileExistsInStorage(relPath)
}

func (sfs *SimpleFileServer) fileExistsInCache(relPath string) (fileExists bool) {
	return sfs.cache.RecordExists(relPath)
}

func (sfs *SimpleFileServer) fileExistsInStorage(relPath string) (fileExists bool, err error) {
	return file.FileExists(sfs.GetAbsolutePath(relPath))
}
