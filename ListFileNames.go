package sfs

import (
	"errors"

	"github.com/vault-thirteen/auxie/file"
)

func (sfs *SimpleFileServer) ListFileNames(folderRelPath string) (fileNames []string, err error) {
	if !IsPathValid(folderRelPath) {
		return nil, errors.New(Err_PathIsNotValid)
	}

	if !sfs.isCachingEnabled {
		return sfs.listFileNamesFromStorage(folderRelPath)
	}

	return sfs.listFileNamesFromStorage(folderRelPath)
}

func (sfs *SimpleFileServer) listFileNamesFromCache(folderRelPath string) (fileNames []string, err error) {
	return nil, errors.New(Err_ActionIsImpossible)
}

func (sfs *SimpleFileServer) listFileNamesFromStorage(folderRelPath string) (fileNames []string, err error) {
	return file.ListFileNames(sfs.GetAbsolutePath(folderRelPath))
}
