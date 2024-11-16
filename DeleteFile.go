package sfs

import (
	"errors"
	"os"
)

func (sfs *SimpleFileServer) DeleteFile(relPath string) (err error) {
	if !IsPathValid(relPath) {
		return errors.New(Err_PathIsNotValid)
	}

	if !sfs.isCachingEnabled {
		return sfs.deleteFileFromStorage(relPath)
	}

	err = sfs.deleteFileFromCache(relPath)
	if err != nil {
		return err
	}

	err = sfs.deleteFileFromStorage(relPath)
	if err != nil {
		return err
	}

	return nil
}

func (sfs *SimpleFileServer) deleteFileFromCache(relPath string) (err error) {
	_, err = sfs.cache.RemoveRecord(relPath)
	return err
}

func (sfs *SimpleFileServer) deleteFileFromStorage(relPath string) (err error) {
	return os.Remove(sfs.GetAbsolutePath(relPath))
}
