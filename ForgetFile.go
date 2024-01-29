package sfs

import (
	"errors"
)

// ForgetFile removes a record about the specified file from cache. Path is a
// relative path inside the file server's root folder.
func (sfs *SimpleFileServer) ForgetFile(relPath string) (err error) {
	if !IsPathValid(relPath) {
		return errors.New(ErrPathIsNotValid)
	}

	_, err = sfs.cache.RemoveRecord(relPath)
	if err != nil {
		return err
	}

	return nil
}
