package sfs

import (
	"errors"
	"path/filepath"
)

// ForgetFile removes a record about the specified file from cache. Path is a
// relative path inside the file server's root folder.
func (sfs *SimpleFileServer) ForgetFile(relPath string) (err error) {
	if !IsPathValid(relPath) {
		return errors.New(ErrPathIsNotValid)
	}

	absFilePath := filepath.Join(sfs.rootFolderPath, relPath)

	_, err = sfs.cache.RemoveRecord(absFilePath)
	if err != nil {
		return err
	}

	return nil
}
