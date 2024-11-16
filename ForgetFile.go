package sfs

import (
	"errors"
)

func (sfs *SimpleFileServer) ForgetFile(relPath string) (err error) {
	if !IsPathValid(relPath) {
		return errors.New(Err_PathIsNotValid)
	}

	if !sfs.isCachingEnabled {
		return errors.New(Err_CacheIsDisabled)
	}

	return sfs.deleteFileFromCache(relPath)
}
