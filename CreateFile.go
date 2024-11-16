package sfs

import (
	"errors"
	"os"
)

func (sfs *SimpleFileServer) CreateFile(relPath string, data []byte) (err error) {
	if !IsPathValid(relPath) {
		return errors.New(Err_PathIsNotValid)
	}

	if !sfs.isCachingEnabled {
		return sfs.createFileInStorage(relPath, data)
	}

	err = sfs.createFileInCache(relPath, data)
	if err != nil {
		return err
	}

	err = sfs.createFileInStorage(relPath, data)
	if err != nil {
		return err
	}

	return nil
}

func (sfs *SimpleFileServer) createFileInCache(relPath string, data []byte) (err error) {
	return sfs.cache.AddRecord(relPath, data)
}

func (sfs *SimpleFileServer) createFileInStorage(relPath string, data []byte) (err error) {
	return os.WriteFile(sfs.GetAbsolutePath(relPath), data, 0777)
}
