package sfs

import (
	"errors"
)

const (
	Err_CacheIsDisabled = "cache is disabled"
)

func (sfs *SimpleFileServer) ForgetFile(relPath string) (err error) {
	if !sfs.isCachingEnabled {
		return errors.New(Err_CacheIsDisabled)
	}

	if !IsPathValid(relPath) {
		return errors.New(ErrPathIsNotValid)
	}

	_, err = sfs.cache.RemoveRecord(relPath)
	if err != nil {
		return err
	}

	return nil
}
