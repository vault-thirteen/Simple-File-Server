package sfs

import (
	"errors"

	"github.com/vault-thirteen/auxie/file"
)

func (sfs *SimpleFileServer) CountFiles(relPath string) (filesCount int, err error) {
	if !IsPathValid(relPath) {
		return -1, errors.New(Err_PathIsNotValid)
	}

	if !sfs.isCachingEnabled {
		return sfs.countFilesInStorage(relPath)
	}

	return sfs.countFilesInStorage(relPath)
}

func (sfs *SimpleFileServer) countFilesInCache(relPath string) (filesCount int, err error) {
	return -1, errors.New(Err_ActionIsImpossible)
}

func (sfs *SimpleFileServer) countFilesInStorage(relPath string) (filesCount int, err error) {
	return file.CountFiles(sfs.GetAbsolutePath(relPath))
}
