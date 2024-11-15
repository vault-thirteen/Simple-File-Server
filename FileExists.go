package sfs

import (
	"path/filepath"

	"github.com/vault-thirteen/auxie/file"
)

func (sfs *SimpleFileServer) FileExists(relPath string) (fileExists bool, err error) {
	if sfs.isCachingEnabled {
		if sfs.fileExistsInCache(relPath) {
			return true, nil
		}
	}

	return sfs.fileExistsInStorage(relPath)
}

func (sfs *SimpleFileServer) fileExistsInCache(relPath string) (fileExists bool) {
	return sfs.cache.RecordExists(relPath)
}

func (sfs *SimpleFileServer) fileExistsInStorage(relPath string) (fileExists bool, err error) {
	absFilePath := filepath.Join(sfs.rootFolderPath, relPath)
	return file.FileExists(absFilePath)
}
