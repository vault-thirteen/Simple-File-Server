package sfs

import (
	"os"
	"path/filepath"
)

// DeleteFile deletes a file from the storage.
func (sfs *SimpleFileServer) DeleteFile(relPath string) (err error) {
	if sfs.isCachingEnabled {
		_, _ = sfs.cache.RemoveRecord(relPath)
	}

	absFilePath := filepath.Join(sfs.rootFolderPath, relPath)

	return os.Remove(absFilePath)
}
