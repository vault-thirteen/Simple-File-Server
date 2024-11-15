package sfs

import (
	"errors"
	"path/filepath"

	"github.com/vault-thirteen/auxie/file"
)

func (sfs *SimpleFileServer) GetFile(relPath string) (bytes []byte, err error) {
	if !IsPathValid(relPath) {
		return nil, errors.New(ErrPathIsNotValid)
	}

	if !sfs.isCachingEnabled {
		return sfs.getFileWithoutCache(relPath)
	}

	if sfs.fileExistsInCache(relPath) {
		return sfs.getFileUsingCache(relPath)
	}

	bytes, err = sfs.getFileWithoutCache(relPath)
	if err != nil {
		return nil, err
	}

	err = sfs.cache.AddRecord(relPath, bytes)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func (sfs *SimpleFileServer) getFileUsingCache(relPath string) (bytes []byte, err error) {
	return sfs.cache.GetRecord(relPath)
}

func (sfs *SimpleFileServer) getFileWithoutCache(relPath string) (bytes []byte, err error) {
	absFilePath := filepath.Join(sfs.rootFolderPath, relPath)

	return file.GetFileContents(absFilePath)
}
