package sfs

import (
	"os"
	"path/filepath"
)

// DeleteFile deletes a file from the storage.
func (sfs *SimpleFileServer) DeleteFile(relPath string) (err error) {
	absFilePath := filepath.Join(sfs.rootFolderPath, relPath)

	err = os.Remove(absFilePath)
	if err != nil {
		return err
	}

	return nil
}
